package app_dbconn

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/client/db"
	"github.com/teejays/gokutil/panics"

	svcauth_dbconn "sampleapp/backend/.goku/generated/service/auth/dbconn"
	svccore_dbconn "sampleapp/backend/.goku/generated/service/core/dbconn"
	svcuser_dbconn "sampleapp/backend/.goku/generated/service/user/dbconn"
)

type AppInitConnectionRequest struct {
	// If specified, this will be used for all services unless specific service connection options are specified for that service.
	Common *db.ServiceInitConnectionRequest
	Auth   *db.ServiceInitConnectionRequest
	Core   *db.ServiceInitConnectionRequest
	User   *db.ServiceInitConnectionRequest
}

// InitDatabaseConnection initializes DB connection for this app.
func InitDatabaseConnection(ctx context.Context, req AppInitConnectionRequest) error {

	var err error

	// Validate the request. If the common options are not specified, then the service options must all be specified.
	if req.Common == nil {
		if req.Auth == nil {
			return fmt.Errorf("Connection options for service must be specified if common options are not specified. Service: [%s]", "Auth")
		}
		if req.Core == nil {
			return fmt.Errorf("Connection options for service must be specified if common options are not specified. Service: [%s]", "Core")
		}
		if req.User == nil {
			return fmt.Errorf("Connection options for service must be specified if common options are not specified. Service: [%s]", "User")
		}
	}
	{
		svcReq := req.Auth
		if svcReq == nil {
			svcReq = req.Common
		}
		// This should never happen since we already validated the request above.
		panics.IfNil(svcReq, "Connection options for service must be specified if common options are not specified. Service: [%s]", "Auth")

		err = svcauth_dbconn.InitDatabaseConnection(ctx, *svcReq)
		if err != nil {
			return fmt.Errorf("Initializing database connection for service [%s]: %w", "Auth", err)
		}
	}
	{
		svcReq := req.Core
		if svcReq == nil {
			svcReq = req.Common
		}
		// This should never happen since we already validated the request above.
		panics.IfNil(svcReq, "Connection options for service must be specified if common options are not specified. Service: [%s]", "Core")

		err = svccore_dbconn.InitDatabaseConnection(ctx, *svcReq)
		if err != nil {
			return fmt.Errorf("Initializing database connection for service [%s]: %w", "Core", err)
		}
	}
	{
		svcReq := req.User
		if svcReq == nil {
			svcReq = req.Common
		}
		// This should never happen since we already validated the request above.
		panics.IfNil(svcReq, "Connection options for service must be specified if common options are not specified. Service: [%s]", "User")

		err = svcuser_dbconn.InitDatabaseConnection(ctx, *svcReq)
		if err != nil {
			return fmt.Errorf("Initializing database connection for service [%s]: %w", "User", err)
		}
	}

	return nil
}
