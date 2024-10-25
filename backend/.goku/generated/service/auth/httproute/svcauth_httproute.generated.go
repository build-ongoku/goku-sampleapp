package svcauth_httproute

import (
	"context"

	gopi "github.com/teejays/gokutil/gopi"
	"github.com/teejays/gokutil/log"

	svcauth_client "sampleapp/backend/.goku/generated/service/auth/client"
	svcauth_entsecret_httproute "sampleapp/backend/.goku/generated/service/auth/entity/secret/httproute"
)

var llog = log.GetLogger().WithHeading("HTTP Handler")

// GetRoutes returns all the routes for this namespace
func GetRoutes(ctx context.Context, c svcauth_client.Client) []gopi.Route {

	var routes []gopi.Route
	// Routes specific to this namespace
	routes = []gopi.Route{
		// Routes at this level
		{
			// API Route for POST /authenticate_token
			Method:       "POST",
			Version:      1,
			Path:         "/auth/authenticate_token",
			HandlerFunc:  gopi.HandlerWrapper("POST", c.AuthenticateToken),
			Authenticate: false,
		},
		{
			// API Route for POST /login
			Method:       "POST",
			Version:      1,
			Path:         "/auth/login",
			HandlerFunc:  gopi.HandlerWrapper("POST", c.LoginUser),
			Authenticate: false,
		},
		{
			// API Route for POST /register
			Method:       "POST",
			Version:      1,
			Path:         "/auth/register",
			HandlerFunc:  gopi.HandlerWrapper("POST", c.RegisterUser),
			Authenticate: false,
		},
	}

	// Routes under this namespace
	routes = append(routes, svcauth_entsecret_httproute.GetRoutes(ctx, c.Secret())...)

	return routes
}
