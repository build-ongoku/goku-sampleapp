package svccore_entsecretkey_httproute

import (
	"context"

	gopi "github.com/teejays/gokutil/gopi"
	"github.com/teejays/gokutil/log"

	svccore_entsecretkey_client "sampleapp/backend/.goku/generated/service/core/entity/secret_key/client"
)

var llog = log.GetLogger().WithHeading("HTTP Handler")

// GetRoutes returns all the routes for this namespace
func GetRoutes(ctx context.Context, c svccore_entsecretkey_client.Client) []gopi.Route {

	var routes []gopi.Route
	// Routes specific to this namespace
	routes = []gopi.Route{
		// Routes at this level
		{
			// API Route for POST
			Method:       "POST",
			Version:      1,
			Path:         "/core/secret_key",
			HandlerFunc:  gopi.HandlerWrapper("POST", c.Add),
			Authenticate: true,
		},
		{
			// API Route for GET
			Method:       "GET",
			Version:      1,
			Path:         "/core/secret_key",
			HandlerFunc:  gopi.HandlerWrapper("GET", c.Get),
			Authenticate: true,
		},
		{
			// API Route for GET /list
			Method:       "GET",
			Version:      1,
			Path:         "/core/secret_key/list",
			HandlerFunc:  gopi.HandlerWrapper("GET", c.List),
			Authenticate: true,
		},
		{
			// API Route for GET /query_by_text
			Method:       "GET",
			Version:      1,
			Path:         "/core/secret_key/query_by_text",
			HandlerFunc:  gopi.HandlerWrapper("GET", c.QueryByText),
			Authenticate: true,
		},
		{
			// API Route for PUT
			Method:       "PUT",
			Version:      1,
			Path:         "/core/secret_key",
			HandlerFunc:  gopi.HandlerWrapper("PUT", c.Update),
			Authenticate: true,
		},
	}

	// Routes under this namespace

	return routes
}
