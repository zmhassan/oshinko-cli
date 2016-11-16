// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_v1_BinaryBuildRequestOptions,
		DeepCopy_v1_BinaryBuildSource,
		DeepCopy_v1_Build,
		DeepCopy_v1_BuildConfig,
		DeepCopy_v1_BuildConfigList,
		DeepCopy_v1_BuildConfigSpec,
		DeepCopy_v1_BuildConfigStatus,
		DeepCopy_v1_BuildList,
		DeepCopy_v1_BuildLog,
		DeepCopy_v1_BuildLogOptions,
		DeepCopy_v1_BuildOutput,
		DeepCopy_v1_BuildPostCommitSpec,
		DeepCopy_v1_BuildRequest,
		DeepCopy_v1_BuildSource,
		DeepCopy_v1_BuildSpec,
		DeepCopy_v1_BuildStatus,
		DeepCopy_v1_BuildStrategy,
		DeepCopy_v1_BuildTriggerPolicy,
		DeepCopy_v1_CustomBuildStrategy,
		DeepCopy_v1_DockerBuildStrategy,
		DeepCopy_v1_GenericWebHookEvent,
		DeepCopy_v1_GitBuildSource,
		DeepCopy_v1_GitInfo,
		DeepCopy_v1_GitSourceRevision,
		DeepCopy_v1_ImageChangeTrigger,
		DeepCopy_v1_ImageSource,
		DeepCopy_v1_ImageSourcePath,
		DeepCopy_v1_JenkinsPipelineBuildStrategy,
		DeepCopy_v1_SecretBuildSource,
		DeepCopy_v1_SecretSpec,
		DeepCopy_v1_SourceBuildStrategy,
		DeepCopy_v1_SourceControlUser,
		DeepCopy_v1_SourceRevision,
		DeepCopy_v1_WebHookTrigger,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1_BinaryBuildRequestOptions(in BinaryBuildRequestOptions, out *BinaryBuildRequestOptions, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api_v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	out.AsFile = in.AsFile
	out.Commit = in.Commit
	out.Message = in.Message
	out.AuthorName = in.AuthorName
	out.AuthorEmail = in.AuthorEmail
	out.CommitterName = in.CommitterName
	out.CommitterEmail = in.CommitterEmail
	return nil
}

func DeepCopy_v1_BinaryBuildSource(in BinaryBuildSource, out *BinaryBuildSource, c *conversion.Cloner) error {
	out.AsFile = in.AsFile
	return nil
}

func DeepCopy_v1_Build(in Build, out *Build, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api_v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_BuildSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_BuildStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1_BuildConfig(in BuildConfig, out *BuildConfig, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api_v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_BuildConfigSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_BuildConfigStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1_BuildConfigList(in BuildConfigList, out *BuildConfigList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]BuildConfig, len(in))
		for i := range in {
			if err := DeepCopy_v1_BuildConfig(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1_BuildConfigSpec(in BuildConfigSpec, out *BuildConfigSpec, c *conversion.Cloner) error {
	if in.Triggers != nil {
		in, out := in.Triggers, &out.Triggers
		*out = make([]BuildTriggerPolicy, len(in))
		for i := range in {
			if err := DeepCopy_v1_BuildTriggerPolicy(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	out.RunPolicy = in.RunPolicy
	if err := DeepCopy_v1_BuildSpec(in.BuildSpec, &out.BuildSpec, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1_BuildConfigStatus(in BuildConfigStatus, out *BuildConfigStatus, c *conversion.Cloner) error {
	out.LastVersion = in.LastVersion
	return nil
}

func DeepCopy_v1_BuildList(in BuildList, out *BuildList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]Build, len(in))
		for i := range in {
			if err := DeepCopy_v1_Build(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1_BuildLog(in BuildLog, out *BuildLog, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1_BuildLogOptions(in BuildLogOptions, out *BuildLogOptions, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Container = in.Container
	out.Follow = in.Follow
	out.Previous = in.Previous
	if in.SinceSeconds != nil {
		in, out := in.SinceSeconds, &out.SinceSeconds
		*out = new(int64)
		**out = *in
	} else {
		out.SinceSeconds = nil
	}
	if in.SinceTime != nil {
		in, out := in.SinceTime, &out.SinceTime
		*out = new(unversioned.Time)
		if err := unversioned.DeepCopy_unversioned_Time(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.SinceTime = nil
	}
	out.Timestamps = in.Timestamps
	if in.TailLines != nil {
		in, out := in.TailLines, &out.TailLines
		*out = new(int64)
		**out = *in
	} else {
		out.TailLines = nil
	}
	if in.LimitBytes != nil {
		in, out := in.LimitBytes, &out.LimitBytes
		*out = new(int64)
		**out = *in
	} else {
		out.LimitBytes = nil
	}
	out.NoWait = in.NoWait
	if in.Version != nil {
		in, out := in.Version, &out.Version
		*out = new(int64)
		**out = *in
	} else {
		out.Version = nil
	}
	return nil
}

func DeepCopy_v1_BuildOutput(in BuildOutput, out *BuildOutput, c *conversion.Cloner) error {
	if in.To != nil {
		in, out := in.To, &out.To
		*out = new(api_v1.ObjectReference)
		if err := api_v1.DeepCopy_v1_ObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.To = nil
	}
	if in.PushSecret != nil {
		in, out := in.PushSecret, &out.PushSecret
		*out = new(api_v1.LocalObjectReference)
		if err := api_v1.DeepCopy_v1_LocalObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.PushSecret = nil
	}
	return nil
}

func DeepCopy_v1_BuildPostCommitSpec(in BuildPostCommitSpec, out *BuildPostCommitSpec, c *conversion.Cloner) error {
	if in.Command != nil {
		in, out := in.Command, &out.Command
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Command = nil
	}
	if in.Args != nil {
		in, out := in.Args, &out.Args
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Args = nil
	}
	out.Script = in.Script
	return nil
}

func DeepCopy_v1_BuildRequest(in BuildRequest, out *BuildRequest, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api_v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Revision != nil {
		in, out := in.Revision, &out.Revision
		*out = new(SourceRevision)
		if err := DeepCopy_v1_SourceRevision(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if in.TriggeredByImage != nil {
		in, out := in.TriggeredByImage, &out.TriggeredByImage
		*out = new(api_v1.ObjectReference)
		if err := api_v1.DeepCopy_v1_ObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.TriggeredByImage = nil
	}
	if in.From != nil {
		in, out := in.From, &out.From
		*out = new(api_v1.ObjectReference)
		if err := api_v1.DeepCopy_v1_ObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	if in.Binary != nil {
		in, out := in.Binary, &out.Binary
		*out = new(BinaryBuildSource)
		if err := DeepCopy_v1_BinaryBuildSource(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Binary = nil
	}
	if in.LastVersion != nil {
		in, out := in.LastVersion, &out.LastVersion
		*out = new(int)
		**out = *in
	} else {
		out.LastVersion = nil
	}
	if in.Env != nil {
		in, out := in.Env, &out.Env
		*out = make([]api_v1.EnvVar, len(in))
		for i := range in {
			if err := api_v1.DeepCopy_v1_EnvVar(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	return nil
}

func DeepCopy_v1_BuildSource(in BuildSource, out *BuildSource, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.Binary != nil {
		in, out := in.Binary, &out.Binary
		*out = new(BinaryBuildSource)
		if err := DeepCopy_v1_BinaryBuildSource(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Binary = nil
	}
	if in.Dockerfile != nil {
		in, out := in.Dockerfile, &out.Dockerfile
		*out = new(string)
		**out = *in
	} else {
		out.Dockerfile = nil
	}
	if in.Git != nil {
		in, out := in.Git, &out.Git
		*out = new(GitBuildSource)
		if err := DeepCopy_v1_GitBuildSource(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Git = nil
	}
	if in.Images != nil {
		in, out := in.Images, &out.Images
		*out = make([]ImageSource, len(in))
		for i := range in {
			if err := DeepCopy_v1_ImageSource(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	out.ContextDir = in.ContextDir
	if in.SourceSecret != nil {
		in, out := in.SourceSecret, &out.SourceSecret
		*out = new(api_v1.LocalObjectReference)
		if err := api_v1.DeepCopy_v1_LocalObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.SourceSecret = nil
	}
	if in.Secrets != nil {
		in, out := in.Secrets, &out.Secrets
		*out = make([]SecretBuildSource, len(in))
		for i := range in {
			if err := DeepCopy_v1_SecretBuildSource(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Secrets = nil
	}
	return nil
}

func DeepCopy_v1_BuildSpec(in BuildSpec, out *BuildSpec, c *conversion.Cloner) error {
	out.ServiceAccount = in.ServiceAccount
	if err := DeepCopy_v1_BuildSource(in.Source, &out.Source, c); err != nil {
		return err
	}
	if in.Revision != nil {
		in, out := in.Revision, &out.Revision
		*out = new(SourceRevision)
		if err := DeepCopy_v1_SourceRevision(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if err := DeepCopy_v1_BuildStrategy(in.Strategy, &out.Strategy, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_BuildOutput(in.Output, &out.Output, c); err != nil {
		return err
	}
	if err := api_v1.DeepCopy_v1_ResourceRequirements(in.Resources, &out.Resources, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_BuildPostCommitSpec(in.PostCommit, &out.PostCommit, c); err != nil {
		return err
	}
	if in.CompletionDeadlineSeconds != nil {
		in, out := in.CompletionDeadlineSeconds, &out.CompletionDeadlineSeconds
		*out = new(int64)
		**out = *in
	} else {
		out.CompletionDeadlineSeconds = nil
	}
	return nil
}

func DeepCopy_v1_BuildStatus(in BuildStatus, out *BuildStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	out.Cancelled = in.Cancelled
	out.Reason = in.Reason
	out.Message = in.Message
	if in.StartTimestamp != nil {
		in, out := in.StartTimestamp, &out.StartTimestamp
		*out = new(unversioned.Time)
		if err := unversioned.DeepCopy_unversioned_Time(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.StartTimestamp = nil
	}
	if in.CompletionTimestamp != nil {
		in, out := in.CompletionTimestamp, &out.CompletionTimestamp
		*out = new(unversioned.Time)
		if err := unversioned.DeepCopy_unversioned_Time(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.CompletionTimestamp = nil
	}
	out.Duration = in.Duration
	out.OutputDockerImageReference = in.OutputDockerImageReference
	if in.Config != nil {
		in, out := in.Config, &out.Config
		*out = new(api_v1.ObjectReference)
		if err := api_v1.DeepCopy_v1_ObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Config = nil
	}
	return nil
}

func DeepCopy_v1_BuildStrategy(in BuildStrategy, out *BuildStrategy, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.DockerStrategy != nil {
		in, out := in.DockerStrategy, &out.DockerStrategy
		*out = new(DockerBuildStrategy)
		if err := DeepCopy_v1_DockerBuildStrategy(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.DockerStrategy = nil
	}
	if in.SourceStrategy != nil {
		in, out := in.SourceStrategy, &out.SourceStrategy
		*out = new(SourceBuildStrategy)
		if err := DeepCopy_v1_SourceBuildStrategy(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.SourceStrategy = nil
	}
	if in.CustomStrategy != nil {
		in, out := in.CustomStrategy, &out.CustomStrategy
		*out = new(CustomBuildStrategy)
		if err := DeepCopy_v1_CustomBuildStrategy(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.CustomStrategy = nil
	}
	if in.JenkinsPipelineStrategy != nil {
		in, out := in.JenkinsPipelineStrategy, &out.JenkinsPipelineStrategy
		*out = new(JenkinsPipelineBuildStrategy)
		if err := DeepCopy_v1_JenkinsPipelineBuildStrategy(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.JenkinsPipelineStrategy = nil
	}
	return nil
}

func DeepCopy_v1_BuildTriggerPolicy(in BuildTriggerPolicy, out *BuildTriggerPolicy, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.GitHubWebHook != nil {
		in, out := in.GitHubWebHook, &out.GitHubWebHook
		*out = new(WebHookTrigger)
		if err := DeepCopy_v1_WebHookTrigger(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.GitHubWebHook = nil
	}
	if in.GenericWebHook != nil {
		in, out := in.GenericWebHook, &out.GenericWebHook
		*out = new(WebHookTrigger)
		if err := DeepCopy_v1_WebHookTrigger(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.GenericWebHook = nil
	}
	if in.ImageChange != nil {
		in, out := in.ImageChange, &out.ImageChange
		*out = new(ImageChangeTrigger)
		if err := DeepCopy_v1_ImageChangeTrigger(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.ImageChange = nil
	}
	return nil
}

func DeepCopy_v1_CustomBuildStrategy(in CustomBuildStrategy, out *CustomBuildStrategy, c *conversion.Cloner) error {
	if err := api_v1.DeepCopy_v1_ObjectReference(in.From, &out.From, c); err != nil {
		return err
	}
	if in.PullSecret != nil {
		in, out := in.PullSecret, &out.PullSecret
		*out = new(api_v1.LocalObjectReference)
		if err := api_v1.DeepCopy_v1_LocalObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		in, out := in.Env, &out.Env
		*out = make([]api_v1.EnvVar, len(in))
		for i := range in {
			if err := api_v1.DeepCopy_v1_EnvVar(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ExposeDockerSocket = in.ExposeDockerSocket
	out.ForcePull = in.ForcePull
	if in.Secrets != nil {
		in, out := in.Secrets, &out.Secrets
		*out = make([]SecretSpec, len(in))
		for i := range in {
			if err := DeepCopy_v1_SecretSpec(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Secrets = nil
	}
	out.BuildAPIVersion = in.BuildAPIVersion
	return nil
}

func DeepCopy_v1_DockerBuildStrategy(in DockerBuildStrategy, out *DockerBuildStrategy, c *conversion.Cloner) error {
	if in.From != nil {
		in, out := in.From, &out.From
		*out = new(api_v1.ObjectReference)
		if err := api_v1.DeepCopy_v1_ObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	if in.PullSecret != nil {
		in, out := in.PullSecret, &out.PullSecret
		*out = new(api_v1.LocalObjectReference)
		if err := api_v1.DeepCopy_v1_LocalObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	out.NoCache = in.NoCache
	if in.Env != nil {
		in, out := in.Env, &out.Env
		*out = make([]api_v1.EnvVar, len(in))
		for i := range in {
			if err := api_v1.DeepCopy_v1_EnvVar(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ForcePull = in.ForcePull
	out.DockerfilePath = in.DockerfilePath
	return nil
}

func DeepCopy_v1_GenericWebHookEvent(in GenericWebHookEvent, out *GenericWebHookEvent, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.Git != nil {
		in, out := in.Git, &out.Git
		*out = new(GitInfo)
		if err := DeepCopy_v1_GitInfo(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Git = nil
	}
	if in.Env != nil {
		in, out := in.Env, &out.Env
		*out = make([]api_v1.EnvVar, len(in))
		for i := range in {
			if err := api_v1.DeepCopy_v1_EnvVar(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	return nil
}

func DeepCopy_v1_GitBuildSource(in GitBuildSource, out *GitBuildSource, c *conversion.Cloner) error {
	out.URI = in.URI
	out.Ref = in.Ref
	if in.HTTPProxy != nil {
		in, out := in.HTTPProxy, &out.HTTPProxy
		*out = new(string)
		**out = *in
	} else {
		out.HTTPProxy = nil
	}
	if in.HTTPSProxy != nil {
		in, out := in.HTTPSProxy, &out.HTTPSProxy
		*out = new(string)
		**out = *in
	} else {
		out.HTTPSProxy = nil
	}
	return nil
}

func DeepCopy_v1_GitInfo(in GitInfo, out *GitInfo, c *conversion.Cloner) error {
	if err := DeepCopy_v1_GitBuildSource(in.GitBuildSource, &out.GitBuildSource, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_GitSourceRevision(in.GitSourceRevision, &out.GitSourceRevision, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1_GitSourceRevision(in GitSourceRevision, out *GitSourceRevision, c *conversion.Cloner) error {
	out.Commit = in.Commit
	if err := DeepCopy_v1_SourceControlUser(in.Author, &out.Author, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_SourceControlUser(in.Committer, &out.Committer, c); err != nil {
		return err
	}
	out.Message = in.Message
	return nil
}

func DeepCopy_v1_ImageChangeTrigger(in ImageChangeTrigger, out *ImageChangeTrigger, c *conversion.Cloner) error {
	out.LastTriggeredImageID = in.LastTriggeredImageID
	if in.From != nil {
		in, out := in.From, &out.From
		*out = new(api_v1.ObjectReference)
		if err := api_v1.DeepCopy_v1_ObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	return nil
}

func DeepCopy_v1_ImageSource(in ImageSource, out *ImageSource, c *conversion.Cloner) error {
	if err := api_v1.DeepCopy_v1_ObjectReference(in.From, &out.From, c); err != nil {
		return err
	}
	if in.Paths != nil {
		in, out := in.Paths, &out.Paths
		*out = make([]ImageSourcePath, len(in))
		for i := range in {
			if err := DeepCopy_v1_ImageSourcePath(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Paths = nil
	}
	if in.PullSecret != nil {
		in, out := in.PullSecret, &out.PullSecret
		*out = new(api_v1.LocalObjectReference)
		if err := api_v1.DeepCopy_v1_LocalObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	return nil
}

func DeepCopy_v1_ImageSourcePath(in ImageSourcePath, out *ImageSourcePath, c *conversion.Cloner) error {
	out.SourcePath = in.SourcePath
	out.DestinationDir = in.DestinationDir
	return nil
}

func DeepCopy_v1_JenkinsPipelineBuildStrategy(in JenkinsPipelineBuildStrategy, out *JenkinsPipelineBuildStrategy, c *conversion.Cloner) error {
	out.JenkinsfilePath = in.JenkinsfilePath
	out.Jenkinsfile = in.Jenkinsfile
	return nil
}

func DeepCopy_v1_SecretBuildSource(in SecretBuildSource, out *SecretBuildSource, c *conversion.Cloner) error {
	if err := api_v1.DeepCopy_v1_LocalObjectReference(in.Secret, &out.Secret, c); err != nil {
		return err
	}
	out.DestinationDir = in.DestinationDir
	return nil
}

func DeepCopy_v1_SecretSpec(in SecretSpec, out *SecretSpec, c *conversion.Cloner) error {
	if err := api_v1.DeepCopy_v1_LocalObjectReference(in.SecretSource, &out.SecretSource, c); err != nil {
		return err
	}
	out.MountPath = in.MountPath
	return nil
}

func DeepCopy_v1_SourceBuildStrategy(in SourceBuildStrategy, out *SourceBuildStrategy, c *conversion.Cloner) error {
	if err := api_v1.DeepCopy_v1_ObjectReference(in.From, &out.From, c); err != nil {
		return err
	}
	if in.PullSecret != nil {
		in, out := in.PullSecret, &out.PullSecret
		*out = new(api_v1.LocalObjectReference)
		if err := api_v1.DeepCopy_v1_LocalObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		in, out := in.Env, &out.Env
		*out = make([]api_v1.EnvVar, len(in))
		for i := range in {
			if err := api_v1.DeepCopy_v1_EnvVar(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.Scripts = in.Scripts
	out.Incremental = in.Incremental
	out.ForcePull = in.ForcePull
	return nil
}

func DeepCopy_v1_SourceControlUser(in SourceControlUser, out *SourceControlUser, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Email = in.Email
	return nil
}

func DeepCopy_v1_SourceRevision(in SourceRevision, out *SourceRevision, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.Git != nil {
		in, out := in.Git, &out.Git
		*out = new(GitSourceRevision)
		if err := DeepCopy_v1_GitSourceRevision(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Git = nil
	}
	return nil
}

func DeepCopy_v1_WebHookTrigger(in WebHookTrigger, out *WebHookTrigger, c *conversion.Cloner) error {
	out.Secret = in.Secret
	out.AllowEnv = in.AllowEnv
	return nil
}
