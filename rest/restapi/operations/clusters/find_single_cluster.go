package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindSingleClusterHandlerFunc turns a function with the right signature into a find single cluster handler
type FindSingleClusterHandlerFunc func(FindSingleClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindSingleClusterHandlerFunc) Handle(params FindSingleClusterParams) middleware.Responder {
	return fn(params)
}

// FindSingleClusterHandler interface for that can handle valid find single cluster params
type FindSingleClusterHandler interface {
	Handle(FindSingleClusterParams) middleware.Responder
}

// NewFindSingleCluster creates a new http.Handler for the find single cluster operation
func NewFindSingleCluster(ctx *middleware.Context, handler FindSingleClusterHandler) *FindSingleCluster {
	return &FindSingleCluster{Context: ctx, Handler: handler}
}

/*FindSingleCluster swagger:route GET /clusters/{name} clusters findSingleCluster

Return detailed information about a single cluster

*/
type FindSingleCluster struct {
	Context *middleware.Context
	Handler FindSingleClusterHandler
}

func (o *FindSingleCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewFindSingleClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
