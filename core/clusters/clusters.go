package clusters

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	oclient "github.com/openshift/origin/pkg/client"
	deployapi "github.com/openshift/origin/pkg/deploy/api"
	ocon "github.com/radanalyticsio/oshinko-cli/core/clusters/containers"
	odc "github.com/radanalyticsio/oshinko-cli/core/clusters/deploymentconfigs"
	opt "github.com/radanalyticsio/oshinko-cli/core/clusters/podtemplates"
	"github.com/radanalyticsio/oshinko-cli/core/clusters/probes"
	ort "github.com/radanalyticsio/oshinko-cli/core/clusters/routes"
	osv "github.com/radanalyticsio/oshinko-cli/core/clusters/services"
	kapi "k8s.io/kubernetes/pkg/api"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/selection"
	"k8s.io/kubernetes/pkg/util/sets"
)


const clusterConfigMsg = "invalid cluster configuration"
const missingConfigMsg = "unable to find spark configuration '%s'"
const findDepConfigMsg = "unable to find deployment configs"
const createDepConfigMsg = "unable to create deployment config '%s'"
const depMsg = "unable to find deployment config for spark %s"
const masterSrvMsg = "unable to create spark master service endpoint"
const masterWebSrvMsg = "unable to create spark master web service endpoint"
const mastermsg = "unable to find spark masters"
const updateDepMsg = "unable to update deployment config for spark %s"
const noSuchClusterMsg = "no such cluster '%s'"
const noClusterForDriverMsg = "no cluster found for app '%s'"
const ephemeralDelMsg = "cluster not deleted '%s'"
const podListMsg = "unable to retrive pod list"
const sparkImageMsg = "no spark image specified"
const clusterExistsMsg = "cluster '%s' already exists%s"

const typeLabel = "oshinko-type"
const clusterLabel = "oshinko-cluster"
const driverLabel = "uses-oshinko-cluster"
const ephemeralLabel = "ephemeral"

const workerType = "worker"
const masterType = "master"
const webuiType = "webui"

const masterPortName = "spark-master"
const masterPort = 7077
const webPortName = "spark-webui"
const webPort = 8080

const sparkconfdir = "/etc/oshinko-spark-configs"

// The suffix to add to the spark master hostname (clustername) for the web service
const webServiceSuffix = "-ui"

type SparkPod struct {
	IP     string
	Status string
	Type   string
}

type SparkCluster struct {
	Namespace    string `json:"namespace,omitempty"`
	Name         string `json:"name,omitempty"`
	Href         string `json:"href"`
	Image        string `json:"image"`
	MasterURL    string `json:"masterUrl"`
	MasterWebURL string `json:"masterWebUrl"`
	MasterWebRoute string `json:"masterWebRoute"`
	Status       string `json:"status"`
	WorkerCount  int    `json:"workerCount"`
	MasterCount  int    `json:"masterCount"`
	Config       ClusterConfig
	Ephemeral    string `json:"ephemeral,omitempty"`
	Pods         []SparkPod
}

func generalErr(err error, msg string, code int) ClusterError {
	if err != nil {
		if msg == "" {
			msg = "error: " + err.Error()
		} else {
			msg = msg + ", error: " + err.Error()
		}
	}
	return NewClusterError(msg, code)
}

func makeSelector(otype string, clustername string) kapi.ListOptions {
	// Build a selector list based on type and/or cluster name
	var ot *labels.Requirement
	var cname *labels.Requirement
	ls := labels.NewSelector()
	if otype == "" {
		ot, _ = labels.NewRequirement(typeLabel, selection.Exists, sets.String{})
	} else {
		ot, _ = labels.NewRequirement(typeLabel, selection.Equals, sets.NewString(otype))
	}
	ls = ls.Add(*ot)
	if clustername == "" {
		cname, _ = labels.NewRequirement(clusterLabel, selection.Exists, sets.String{})
	} else {
		cname, _ = labels.NewRequirement(clusterLabel, selection.Equals, sets.NewString(clustername))
	}
	ls = ls.Add(*cname)
	return kapi.ListOptions{LabelSelector: ls}
}

func retrieveServiceURL(client kclient.ServiceInterface, stype, clustername string) string {
	selectorlist := makeSelector(stype, clustername)
	srvs, err := client.List(selectorlist)
	if err == nil && len(srvs.Items) != 0 {
		srv := srvs.Items[0]
		scheme := "http://"
		if stype == masterType {
			scheme = "spark://"
		}
		return scheme + srv.Name + ":" + strconv.Itoa(int(srv.Spec.Ports[0].Port))
	}
	return "<missing>"
}

func retrieveRouteForService(client oclient.RouteInterface, stype, clustername string) string {
	selectorlist := makeSelector(stype, clustername)
	routes, err := client.List(selectorlist)
	if err == nil && len(routes.Items) != 0 {
		route := routes.Items[0]
		return route.Spec.Host
	}
	return "<no route>"
}

func checkForDeploymentConfigs(client oclient.DeploymentConfigInterface, clustername string) (bool, *deployapi.DeploymentConfig, error) {
	selectorlist := makeSelector(masterType, clustername)
	dcs, err := client.List(selectorlist)
	if err != nil {
		return false, nil, err
	}
	if len(dcs.Items) == 0 {
		return false, nil, nil
	}
	m := dcs.Items[0]
	selectorlist = makeSelector(workerType, clustername)
	dcs, err = client.List(selectorlist)
	if err != nil {
		return false, &m, err
	}
	if len(dcs.Items) == 0 {
		return false, &m, nil
	}
	return true, &m, nil

}

func makeEnvVars(clustername, sparkconfdir string) []kapi.EnvVar {
	envs := []kapi.EnvVar{}

	envs = append(envs, kapi.EnvVar{Name: "OSHINKO_SPARK_CLUSTER", Value: clustername})
	envs = append(envs, kapi.EnvVar{Name: "OSHINKO_REST_HOST", Value: os.Getenv("OSHINKO_REST_SERVICE_HOST")})
	envs = append(envs, kapi.EnvVar{Name: "OSHINKO_REST_PORT", Value: os.Getenv("OSHINKO_REST_SERVICE_PORT")})
	if sparkconfdir != "" {
		envs = append(envs, kapi.EnvVar{Name: "UPDATE_SPARK_CONF_DIR", Value: sparkconfdir})
	}

	return envs
}

func makeWorkerEnvVars(clustername, sparkconfdir string) []kapi.EnvVar {
	envs := []kapi.EnvVar{}

	envs = makeEnvVars(clustername, sparkconfdir)
	envs = append(envs, kapi.EnvVar{
		Name:  "SPARK_MASTER_ADDRESS",
		Value: "spark://" + clustername + ":" + strconv.Itoa(masterPort)})
	envs = append(envs, kapi.EnvVar{
		Name:  "SPARK_MASTER_UI_ADDRESS",
		Value: "http://" + clustername + webServiceSuffix + ":" + strconv.Itoa(webPort)})
	return envs
}

func sparkWorker(namespace, image string, replicas int, clustername, sparkconfdir, sparkworkerconfig string) *odc.ODeploymentConfig {

	// Create the basic deployment config
	// We will use a label and pod selector based on the cluster name.
	// Openshift will add additional labels and selectors to distinguish pods handled by
	// this deploymentconfig from pods beloning to another.
	dc := odc.DeploymentConfig(clustername+"-w", namespace).
		TriggerOnConfigChange().RollingStrategy().Label(clusterLabel, clustername).
		Label(typeLabel, workerType).
		PodSelector(clusterLabel, clustername).Replicas(replicas)

	// Create a pod template spec with the matching label
	pt := opt.PodTemplateSpec().Label(clusterLabel, clustername).Label(typeLabel, workerType)

	// Create a container with the correct ports and start command
	webport := 8081
	webp := ocon.ContainerPort(webPortName, webport)
	cont := ocon.Container(dc.Name, image).
		Ports(webp).
		SetLivenessProbe(probes.NewHTTPGetProbe(webport)).EnvVars(makeWorkerEnvVars(clustername, sparkconfdir))

	if sparkworkerconfig != "" {
		pt = pt.SetConfigMapVolume(sparkworkerconfig)
		cont = cont.SetVolumeMount(sparkworkerconfig, sparkconfdir, true)
	}

	// Finally, assign the container to the pod template spec and
	// assign the pod template spec to the deployment config
	return dc.PodTemplateSpec(pt.Containers(cont))
}

func mastername(clustername string) string {
	return clustername + "-m"
}

func workername(clustername string) string {
	return clustername + "-w"
}

func sparkMaster(namespace, image string, replicas int, clustername, sparkconfdir, sparkmasterconfig, driverrc string) *odc.ODeploymentConfig {

	// Create the basic deployment config
	// We will use a label and pod selector based on the cluster name
	// Openshift will add additional labels and selectors to distinguish pods handled by
	// this deploymentconfig from pods beloning to another.
	dc := odc.DeploymentConfig(mastername(clustername), namespace).
		TriggerOnConfigChange().RollingStrategy().Label(clusterLabel, clustername).
		Label(typeLabel, masterType).
		PodSelector(clusterLabel, clustername).Replicas(replicas)

	if driverrc != "" {
		dc = dc.Label(ephemeralLabel, driverrc)
	}

	// Create a pod template spec with the matching label
	pt := opt.PodTemplateSpec().Label(clusterLabel, clustername).
		Label(typeLabel, masterType)

	// Create a container with the correct ports and start command
	liveness := probes.NewExecProbe([]string{"/bin/bash", "-c", "curl localhost:8080 | grep -e Status.*ALIVE"})
	liveness.InitialDelaySeconds = 10
	readiness := probes.NewHTTPGetProbe(webPort)
	masterp := ocon.ContainerPort(masterPortName, masterPort)
	webp := ocon.ContainerPort(webPortName, webPort)
	cont := ocon.Container(dc.Name, image).
		Ports(masterp, webp).
		SetLivenessProbe(liveness).
		SetReadinessProbe(readiness).EnvVars(makeEnvVars(clustername, sparkconfdir))

	if sparkmasterconfig != "" {
		pt = pt.SetConfigMapVolume(sparkmasterconfig)
		cont = cont.SetVolumeMount(sparkmasterconfig, sparkconfdir, true)
	}

	// Finally, assign the container to the pod template spec and
	// assign the pod template spec to the deployment config
	return dc.PodTemplateSpec(pt.Containers(cont))
}

func service(name string,
	port int,
	clustername, otype string,
	podselectors map[string]string) (*osv.OService, *osv.OServicePort) {

	p := osv.ServicePort(port).TargetPort(port)
	return osv.Service(name).Label(clusterLabel, clustername).
		Label(typeLabel, otype).PodSelectors(podselectors).Ports(p), p
}

func checkForConfigMap(name string, cm kclient.ConfigMapsInterface) error {
	_, err := cm.Get(name)
	if err != nil {
		if strings.Index(err.Error(), "not found") != -1 {
			return generalErr(err, fmt.Sprintf(missingConfigMsg, name), ClusterConfigCode)
		}
		return generalErr(nil, fmt.Sprintf(missingConfigMsg, name), ClientOperationCode)
	}
	return nil
}

func countWorkers(client kclient.PodInterface, clustername string) (int, *kapi.PodList, error) {
	// If we are  unable to retrieve a list of worker pods, return -1 for count
	// This is an error case, differnt from a list of length 0. Let the caller
	// decide whether to report the error or the -1 count
	cnt := -1
	selectorlist := makeSelector(workerType, clustername)
	pods, err := client.List(selectorlist)
	if pods != nil {
		cnt = len(pods.Items)
	}
	return cnt, pods, err
}

func getDriverDeployment(app, namespace string, client kclient.Interface) string {

	// When we make calls from a driver pod, the most likely value we have is a deployment
	// so use that first
	rcc := client.ReplicationControllers(namespace)
	rc, err := rcc.Get(app)
	if err == nil && rc != nil {
		return app
	}

	// Okay, it wasn't a deployment, maybe it's a pod
	pc := client.Pods(namespace)
	pod, err := pc.Get(app)
	if err == nil && pod != nil {
		return pod.Labels["deployment"]
	}
	return ""
}

// Create a cluster and return the representation
func CreateCluster(
	clustername, namespace, sparkimage string,
	config *ClusterConfig, osclient *oclient.Client, client kclient.Interface, app string, ephemeral bool) (SparkCluster, error) {

	var driverrc string
	var ephem_val string
	var masterconfdir string
	var workerconfdir string
	var result SparkCluster = SparkCluster{}

	createCode := func(err error) int {
		if err != nil && strings.Index(err.Error(), "already exists") != -1 {
			return ComponentExistsCode
		}
		return ClientOperationCode
	}

	masterhost := clustername

	// Check to see if a cluster already exists of the same name (complete or incomplete)
	existing := SparkCluster{}
	findClusterBody(clustername, namespace, osclient, client, &existing)
	if !CheckNoCluster(&existing) {
		var msg string
		if existing.Status != "Running" {
			msg = fmt.Sprintf(clusterExistsMsg, clustername, " (incomplete)")
		} else {
			msg = fmt.Sprintf(clusterExistsMsg, clustername, "")
		}
		return result, generalErr(nil, msg, ComponentExistsCode)
	}


	// Copy any named config referenced and update it with any explicit config values
	finalconfig, err := GetClusterConfig(config, client.ConfigMaps(namespace))
	if err != nil {
		return result, generalErr(err, clusterConfigMsg, ErrorCode(err))
	}
	if finalconfig.SparkImage != "" {
		sparkimage = finalconfig.SparkImage
	} else if sparkimage == "" {
		return result, generalErr(nil, sparkImageMsg, ClusterConfigCode)
	}

	mastercount := int(finalconfig.MasterCount)
	workercount := int(finalconfig.WorkerCount)

	// Check if finalconfig contains the names of ConfigMaps to use for spark
	// configuration. If so they must exist. The ConfigMaps will be mounted
	// as volumes on spark pods and the path stored in the environment
	// variable UPDATE_SPARK_CONF_DIR
	cm := client.ConfigMaps(namespace)
	if finalconfig.SparkMasterConfig != "" {
		err := checkForConfigMap(finalconfig.SparkMasterConfig, cm)
		if err != nil {
			return result, err
		}
		masterconfdir = sparkconfdir
	}

	if finalconfig.SparkWorkerConfig != "" {
		err := checkForConfigMap(finalconfig.SparkWorkerConfig, cm)
		if err != nil {
			return result, err
		}
		workerconfdir = sparkconfdir
	}

	// If an app value was passed find the deployment so we can do the cluster
	// association and potentially mark the cluster as ephemeral
	if app != "" {
		driverrc = getDriverDeployment(app, namespace, client)
	}
	if ephemeral {
		// If we can't find an rc then we just make a long-running cluster
		ephem_val = driverrc
	}

	// Create the master deployment config
	masterdc := sparkMaster(namespace, sparkimage, mastercount, clustername, masterconfdir, finalconfig.SparkMasterConfig, ephem_val)

	// Create the services that will be associated with the master pod
	// They will be created with selectors based on the pod labels
	mastersv, _ := service(masterhost,
		masterdc.FindPort(masterPortName),
		clustername, masterType,
		masterdc.GetPodTemplateSpecLabels())

	websv, _ := service(masterhost+webServiceSuffix,
		masterdc.FindPort(webPortName),
		clustername, webuiType,
		masterdc.GetPodTemplateSpecLabels())

	webuiroute := ort.NewRoute(websv.GetName() + "-route", websv.GetName(), clustername, "webui")

	// Create the worker deployment config
	workerdc := sparkWorker(namespace, sparkimage, workercount, clustername, workerconfdir, finalconfig.SparkWorkerConfig)

	// Launch all of the objects
	dcc := osclient.DeploymentConfigs(namespace)
	_, err = dcc.Create(&masterdc.DeploymentConfig)
	if err != nil {
		return result, generalErr(err, fmt.Sprintf(createDepConfigMsg, masterdc.Name), createCode(err))
	}
	_, err = dcc.Create(&workerdc.DeploymentConfig)
	if err != nil {
		// Since we created the master deployment config, try to clean up
		DeleteCluster(clustername, namespace, osclient, client, "", "")
		return result, generalErr(err, fmt.Sprintf(createDepConfigMsg, workerdc.Name), createCode(err))
	}

	sc := client.Services(namespace)
	_, err = sc.Create(&mastersv.Service)
	if err != nil {
		// Since we've created things, try to clean up
		DeleteCluster(clustername, namespace, osclient, client, "", "")
		return result, generalErr(err, masterSrvMsg, createCode(err))
	}


	_, err = sc.Create(&websv.Service)
	if err != nil {
		// Since we've created things, try to clean up
		DeleteCluster(clustername, namespace, osclient, client, "", "")
		return result, generalErr(err, masterWebSrvMsg, createCode(err))
	}

	// We will expose the Spark master webui unless we are told not to do it
	if config.ExposeWebUI {
		rc := osclient.Routes(namespace)
		_, err = rc.Create(webuiroute)
	}

	// Wait for the replication controllers to exist before building the response.
	rcc := client.ReplicationControllers(namespace)
	{
		var mrepl, wrepl *kapi.ReplicationController
		mrepl = nil
		wrepl = nil
		for i := 0; i < 4; i++ {
			if mrepl == nil {
				mrepl, _ = getReplController(rcc, clustername, masterType)
			}
			if wrepl == nil {
				wrepl, _ = getReplController(rcc, clustername, workerType)
			}
			if wrepl != nil && mrepl != nil {
				break
			}
			time.Sleep(250 * time.Millisecond)
		}
	}

	// Now that the creation actually worked, label the dc if the app value was passed.
	// Note that updates can fail if someone updates the object underneath us, so
	// we have to try again.  Try for 5 seconds
	if driverrc != "" {
		for i := 0; i < 20; i++ {
			driver, err := rcc.Get(driverrc)
			if err == nil {
				if driver.Labels == nil {
					driver.Labels = map[string]string{}
				}
				driver.Labels[driverLabel] = clustername
				_, err = rcc.Update(driver)
				if err == nil {
					break
				}
				time.Sleep(250 * time.Millisecond)
			} else {
				break
			}
		}
	}

	result.Name = clustername
	result.Namespace = namespace
	result.Href = "/clusters/" + clustername
	result.Image = sparkimage
	result.MasterURL = retrieveServiceURL(sc, masterType, clustername)
	result.MasterWebURL = retrieveServiceURL(sc, webuiType, clustername)
	result.Status = "Running"
	result.Config = finalconfig
	result.MasterCount = 1
	result.WorkerCount = workercount
	result.Pods = []SparkPod{}
	if ephem_val != "" {
		result.Ephemeral = ephem_val
	} else {
		result.Ephemeral = "<shared>"
	}

	return result, nil
}

func DeleteCluster(clustername, namespace string, osclient *oclient.Client, client kclient.Interface, app, appstatus string) (string, error) {
	var foundSomething bool = false
	//var zero int32 = 0
	info := []string{}
	rcnames := []string{}


	dcc := osclient.DeploymentConfigs(namespace)
	rcc := client.ReplicationControllers(namespace)

	// If we have supplied an appstatus flag, then we only delete the cluster if it is marked as ephemeral
	// If it's not marked as ephemeral then we skip the delete
	if appstatus == "completed" || appstatus == "terminated" {
		var delete bool = false

		// See if the master dc has the ephemeral label
		master, err := dcc.Get(mastername(clustername))
		if err != nil {
			// We can't get the dc for the master to look up whether it's ephemeral.
			// But this means the cluster is partially broken anyway. Let the normal delete
			// fall through and cleanup
			delete = true
		} else if ephemeral, ok := master.Labels[ephemeralLabel]; ok {
			deployment := getDriverDeployment(app, namespace, client)
			if deployment != ephemeral {
				info = append(info, "cluster is not linked to app")
			} else {
				// If the driver has been scaled to zero, or if the application
				// completed and the repl count is 1 then delete (because in the
				// completed case the driver is the only instance)
				repl, err := rcc.Get(deployment)
				delete = err != nil ||
					repl.Spec.Replicas == 0 ||
					(appstatus == "completed" && repl.Spec.Replicas == 1)
				if !delete {
					info = append(info, "driver replica count > 0 (or > 1 for completed app)")
				}
			}

		} else {
			info = append(info, "cluster is not ephemeral")
		}
		if !delete {
			return strings.Join(info, ", "), generalErr(nil, fmt.Sprintf(ephemeralDelMsg, clustername), EphemeralCode)
		}
	}

	// Build a selector list for the "oshinko-cluster" label
	selectorlist := makeSelector("", clustername)

	// Delete the dcs
	deployments, err := dcc.List(selectorlist)
	for i := range deployments.Items {
		err = dcc.Delete(deployments.Items[i].Name)
		if err != nil {
			info = append(info, "unable to delete deployment config "+deployments.Items[i].Name+" ("+err.Error()+")")
		} else {
			foundSomething = true
		}
	}

	// Delete the rcs
	rcc = client.ReplicationControllers(namespace)
	repls, err := rcc.List(selectorlist)
	for i := range repls.Items {
		rcnames = append(rcnames, repls.Items[i].Name)
		err = rcc.Delete(repls.Items[i].Name, nil)
		if err != nil {
			info = append(info, "unable to delete replication controller " + repls.Items[i].Name + " (" + err.Error() + ")")
		} else {
			foundSomething = true
		}
	}

	pc := client.Pods(namespace)
	for i := range rcnames {
		ls := labels.NewSelector()
		plist, _ := labels.NewRequirement("openshift.io/deployer-pod-for.name", selection.Equals, sets.NewString(rcnames[i]))
		ls = ls.Add(*plist)
		pods, err := pc.List(kapi.ListOptions{LabelSelector: ls})
		if err == nil && len(pods.Items) != 0 {
			for p := range pods.Items {
				err = pc.Delete(pods.Items[p].Name, nil)
				if err != nil {
					info = append(info, "unable to delete deployer pod " + pods.Items[p].Name + " ("+err.Error()+")")
				} else {
					info = append(info, "deleted deployer pod " + pods.Items[p].Name)
					foundSomething = true
				}
			}
		}
	}

	pods, err := pc.List(selectorlist)
	for i := range pods.Items {
		err = pc.Delete(pods.Items[i].Name, nil)
		if err != nil {
			info = append(info, "unable to delete pod " + pods.Items[i].Name + " (" + err.Error() + ")")
		} else {
			foundSomething = true
		}
	}

	rc := osclient.Routes(namespace)
	webUIRouteName := clustername + "-ui-route"
	err = rc.Delete(webUIRouteName)
	if err != nil {
		info = append(info, "unable to delete route " + webUIRouteName + " (" + err.Error() + ")")
	}

	// Delete the services
	sc := client.Services(namespace)
	srvs, err := sc.List(selectorlist)
	for i := range srvs.Items {
		err = sc.Delete(srvs.Items[i].Name)
		if err != nil {
			info = append(info, "unable to delete service " + srvs.Items[i].Name + " (" + err.Error() + ")")
		} else {
			foundSomething = true
		}
	}
	// If we found some part of a cluster, then there is no error
	// even though the cluster may not have been fully complete.
	// If we didn't find any trace of a cluster, then call it an error
	if !foundSomething {
		return strings.Join(info, ", "), generalErr(nil, fmt.Sprintf(noSuchClusterMsg, clustername), NoSuchClusterCode)
	}
	return strings.Join(info, ", "), nil
}

func findClusterBody(clustername, namespace string, osclient *oclient.Client, client kclient.Interface, result *SparkCluster) {

	sc := client.Services(namespace)
	dc := osclient.DeploymentConfigs(namespace)

	result.Name = clustername
	result.Namespace = namespace
	result.Href = "/clusters/" + clustername

	// TODO make something real for status
	result.Status = "Running"

	// Note, we do not report an error here since we are
	// reporting on multiple clusters. Instead cnt will be 0.
	worker, err := dc.Get(workername(clustername))
	if err == nil {
		result.WorkerCount = int(worker.Status.Replicas)
		result.Config.WorkerCount = int(worker.Spec.Replicas)
	} else {
		result.WorkerCount = -1
		result.Config.WorkerCount = -1
		result.Status = "Incomplete"
	}

	// TODO we only want to count running pods (not terminating)
	result.MasterURL = retrieveServiceURL(sc, masterType, clustername)
	if result.MasterURL == "<missing>" {
		result.Status = "Incomplete"
	}
	result.MasterWebURL = retrieveServiceURL(sc, webuiType, clustername)
	if result.MasterWebURL == "<missing>" {
		result.Status = "Incomplete"
	}
	result.MasterWebRoute = retrieveRouteForService(osclient.Routes(namespace), webuiType, clustername)

	result.Ephemeral = "<shared>"
	master, err := dc.Get(mastername(clustername))
	if err == nil {
		result.MasterCount = int(master.Status.Replicas)
		result.Config.MasterCount = int(master.Spec.Replicas)
		if ephemeral, ok := master.Labels[ephemeralLabel]; ok {
			result.Ephemeral = ephemeral
		}
	} else {
		result.MasterCount = -1
		result.Config.MasterCount = -1
		result.Status = "Incomplete"
	}
}

func CheckNoCluster(cluster *SparkCluster) bool {
	// negative counts here means that there was no dc
	// we might still have pods but they should be terminating, and even if one is stuck, it bears
	// a random suffix so not really a problem
	return cluster.Status == "Incomplete" && cluster.WorkerCount == -1 && cluster.MasterCount == -1 &&
	       cluster.MasterURL == "<missing>" && cluster.MasterWebURL == "<missing>"
}

// Find a cluster and return its representation
func FindSingleCluster(name, namespace string, osclient *oclient.Client, client kclient.Interface) (SparkCluster, error) {

	addpod := func(p kapi.Pod) SparkPod {
		return SparkPod{
			IP:     p.Status.PodIP,
			Status: string(p.Status.Phase),
			Type:   p.Labels[typeLabel],
		}
	}
        var result SparkCluster
	findClusterBody(name, namespace, osclient, client, &result)

	pc := client.Pods(namespace)
	pods, err := pc.List(makeSelector("", name))
	if err != nil {
		return result, generalErr(err, podListMsg, ClientOperationCode)
	}

	// Report pods
	result.Pods = []SparkPod{}
	for i := range pods.Items {
		result.Pods = append(result.Pods, addpod(pods.Items[i]))
	}
	if CheckNoCluster(&result) {
		return result, generalErr(nil, fmt.Sprintf(noSuchClusterMsg, name), NoSuchClusterCode)
	}
	return result, nil
}

// Find all clusters and return their representation
func FindClusters(namespace string, osclient *oclient.Client, client kclient.Interface, app string) ([]SparkCluster, error) {

	dcc := osclient.DeploymentConfigs(namespace)
	rc := client.ReplicationControllers(namespace)

	// If app is not null, look for a driver label.
	// If we find it get the name of the cluster and call FindSingleCluster.
	if app != "" {
		deployment := getDriverDeployment(app, namespace, client)
		if deployment != "" {
			driver, err := rc.Get(deployment)
			if err == nil && driver != nil {
				if clustername, ok := driver.Labels[driverLabel]; ok {
					result := make([]SparkCluster, 1, 1)
					result[0], err = FindSingleCluster(clustername, namespace, osclient, client)
					return result, err
				}
			}
		}
		return []SparkCluster{}, generalErr(nil, fmt.Sprintf(noClusterForDriverMsg, app), NoSuchClusterCode)
	}

	// Create a map so that we can track clusters by name while we
	// find out information about them
	clist := map[string]bool{}

	// Get all of the master dcs
	dcs, err := dcc.List(makeSelector(masterType, ""))
	if err != nil {
		return []SparkCluster{}, generalErr(err, mastermsg, ClientOperationCode)
	}

	// From the list of master dcs, figure out which clusters we have
	for i := range dcs.Items {

		// Add the cluster name if we don't already have it
		clustername := dcs.Items[i].Labels[clusterLabel]
		if _, ok := clist[clustername]; !ok {
			clist[clustername] = true
		}
	}
	result := make([]SparkCluster, len(clist), len(clist))
	idx := 0
	for k, _ := range clist {
		findClusterBody(k, namespace, osclient, client, &result[idx])
		idx++
	}
	return result, nil
}

func newestRepl(list *kapi.ReplicationControllerList ) *kapi.ReplicationController {
	newestRepl := list.Items[0]
	for i := 0; i < len(list.Items); i++ {
		if list.Items[i].CreationTimestamp.Unix() > newestRepl.CreationTimestamp.Unix() {
			newestRepl = list.Items[i]
		}
	}
	return &newestRepl
}

func getReplController(client kclient.ReplicationControllerInterface, clustername, otype string) (*kapi.ReplicationController, error) {

	selectorlist := makeSelector(otype, clustername)
	repls, err := client.List(selectorlist)
	if err != nil || len(repls.Items) == 0 {
		return nil, err
	}
	// Use the latest replication controller.  There could be more than one
	// if the user did something like oc env to set a new env var on a deployment
	return newestRepl(repls), nil
}

func getDepConfig(client oclient.DeploymentConfigInterface, clustername, otype string) (*deployapi.DeploymentConfig, error) {
	var dep *deployapi.DeploymentConfig
	var err error
	if otype == masterType {
		dep, err = client.Get(mastername(clustername))
	} else {
		dep, err = client.Get(workername(clustername))
	}
	return dep, err
}

func scaleDep(dcc oclient.DeploymentConfigInterface, clustername string, count int, otype string) error {
	var err error
	if count <= SentinelCountValue {
		return nil
	}

	for i := 0; i < 20; i++ {
		dep, err := getDepConfig(dcc, clustername, otype)
		if err != nil {
			return generalErr(err, fmt.Sprintf(depMsg, otype), ClientOperationCode)
		} else if dep == nil {
			return generalErr(err, fmt.Sprintf(depMsg, otype), ClusterIncompleteCode)
		}
		// If the current replica count does not match the request, update the replication controller
		if int(dep.Spec.Replicas) != count {
			dep.Spec.Replicas = int32(count)
			_, err = dcc.Update(dep)
			if err == nil {
				break
			}
		} else {
			break
		}
		time.Sleep(250*time.Millisecond)
	}

	// if err has a value here then we failed all retries and err is the last message from a failed update
	if err != nil {
		return generalErr(err, fmt.Sprintf(updateDepMsg, otype), ClientOperationCode)
	}
	return nil
}

// Update a cluster and return the new representation
// This routine supports the same stored config semantics as used in cluster creation
// but at this point only allows updating the master and worker counts.
func UpdateCluster(name, namespace string, config *ClusterConfig, osclient *oclient.Client, client kclient.Interface) (SparkCluster, error) {

	var result SparkCluster = SparkCluster{}
	clustername := name

	// Before we do further checks, make sure that we have deploymentconfigs
	// If either the master or the worker deploymentconfig are missing, we
	// assume that the cluster is missing. These are the base objects that
	// we use to create a cluster
	ok, _, err := checkForDeploymentConfigs(osclient.DeploymentConfigs(namespace), clustername)
	if err != nil {
		return result, generalErr(err, findDepConfigMsg, ClientOperationCode)
	}
	if !ok {
		return result, generalErr(nil, fmt.Sprintf(noSuchClusterMsg, clustername), NoSuchClusterCode)
	}

	// Copy any named config referenced and update it with any explicit config values
	finalconfig, err := GetClusterConfig(config, client.ConfigMaps(namespace))
	if err != nil {
		return result, generalErr(err, clusterConfigMsg, ErrorCode(err))
	}
	workercount := int(finalconfig.WorkerCount)
	mastercount := int(finalconfig.MasterCount)

	// TODO(tmckay) we need some way to track the current spark config for a cluster,
	// maybe in annotations. If someone tries to change the spark config for a cluster,
	// that should be an error at this point (unless we spin all the pods down and
	// redeploy)

	dcc := osclient.DeploymentConfigs(namespace)
	err = scaleDep(dcc, clustername, workercount, workerType)
	if err != nil {
		return result, err
	}
	err = scaleDep(dcc, clustername, mastercount, masterType)
	if err != nil {
		return result, err
	}

	result.Name = name
	result.Namespace = namespace
	result.Config = finalconfig
	return result, nil
}


// Scale a cluster
// This routine supports a specific scale operation based on immediate values for
// master and worker counts and does not consider stored configs.
func ScaleCluster(name, namespace string, masters, workers int, osclient *oclient.Client, client kclient.Interface) error {

	clustername := name

	// Before we do further checks, make sure that we have deploymentconfigs
	// If either the master or the worker deploymentconfig are missing, we
	// assume that the cluster is missing. These are the base objects that
	// we use to create a cluster
	ok, _, err := checkForDeploymentConfigs(osclient.DeploymentConfigs(namespace), clustername)
	if err != nil {
		return generalErr(err, findDepConfigMsg, ClientOperationCode)
	}
	if !ok {
		return generalErr(nil, fmt.Sprintf(noSuchClusterMsg, clustername), NoSuchClusterCode)
	}

	// Allow sale to zero, allow sentinel values
	if masters > 1 {
		return NewClusterError(MasterCountMustBeZeroOrOne, ClusterConfigCode)
	}

	dcc := osclient.DeploymentConfigs(namespace)
	err = scaleDep(dcc, clustername, workers, workerType)
	if err != nil {
		return err
	}
	return scaleDep(dcc, clustername, masters, masterType)
}
