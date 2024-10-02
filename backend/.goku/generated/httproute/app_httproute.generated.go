package app_httproute

import (
	"context"

	gopi "github.com/teejays/gokutil/gopi"
	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_httproute "sampleapp/backend/.goku/generated/service/auth/httproute"
	svccore_httproute "sampleapp/backend/.goku/generated/service/core/httproute"
	svcuser_httproute "sampleapp/backend/.goku/generated/service/user/httproute"
)

var llog = log.GetLogger().WithHeading("HTTP Handler")

// GetRoutes returns all the routes for this namespace
func GetRoutes(ctx context.Context, c app_client.Client) []gopi.Route {

	var routes []gopi.Route

	// Routes under this namespace
	routes = append(routes, svcauth_httproute.GetRoutes(ctx, c.Auth())...)
	routes = append(routes, svccore_httproute.GetRoutes(ctx, c.Core())...)
	routes = append(routes, svcuser_httproute.GetRoutes(ctx, c.User())...)

	return routes
}
