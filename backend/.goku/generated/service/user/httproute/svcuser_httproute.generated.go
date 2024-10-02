package svcuser_httproute

import (
	"context"

	gopi "github.com/teejays/gokutil/gopi"
	"github.com/teejays/gokutil/log"

	svcuser_client "sampleapp/backend/.goku/generated/service/user/client"
	svcuser_entteam_httproute "sampleapp/backend/.goku/generated/service/user/entity/team/httproute"
	svcuser_entuser_httproute "sampleapp/backend/.goku/generated/service/user/entity/user/httproute"
)

var llog = log.GetLogger().WithHeading("HTTP Handler")

// GetRoutes returns all the routes for this namespace
func GetRoutes(ctx context.Context, c svcuser_client.Client) []gopi.Route {

	var routes []gopi.Route
	// Routes specific to this namespace
	routes = []gopi.Route{
		// Routes at this level
	}

	// Routes under this namespace
	routes = append(routes, svcuser_entteam_httproute.GetRoutes(ctx, c.Team())...)
	routes = append(routes, svcuser_entuser_httproute.GetRoutes(ctx, c.User())...)

	return routes
}
