package svccore_httproute

import (
	"context"

	gopi "github.com/teejays/gokutil/gopi"
	"github.com/teejays/gokutil/log"

	svccore_client "sampleapp/backend/.goku/generated/service/core/client"
	svccore_entfile_httproute "sampleapp/backend/.goku/generated/service/core/entity/file/httproute"
	svccore_enttask_httproute "sampleapp/backend/.goku/generated/service/core/entity/task/httproute"
	svccore_enttaskrun_httproute "sampleapp/backend/.goku/generated/service/core/entity/task_run/httproute"
)

var llog = log.GetLogger().WithHeading("HTTP Handler")

// GetRoutes returns all the routes for this namespace
func GetRoutes(ctx context.Context, c svccore_client.Client) []gopi.Route {

	var routes []gopi.Route
	// Routes specific to this namespace
	routes = []gopi.Route{
		// Routes at this level
		{
			// API Route for POST file/upload
			Method:       "POST",
			Version:      1,
			Path:         "/core/file/upload",
			HandlerFunc:  gopi.GetDirectRequestHandler(c.FileUpload),
			Authenticate: true,
		},
	}

	// Routes under this namespace
	routes = append(routes, svccore_entfile_httproute.GetRoutes(ctx, c.File())...)
	routes = append(routes, svccore_enttask_httproute.GetRoutes(ctx, c.Task())...)
	routes = append(routes, svccore_enttaskrun_httproute.GetRoutes(ctx, c.TaskRun())...)

	return routes
}
