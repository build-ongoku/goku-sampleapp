package app_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	app_clientimpl "sampleapp/backend/.goku/generated/clientimpl"
	svcauth_mthd "sampleapp/backend/.goku/generated/service/auth/mthd"
	svccore_mthd "sampleapp/backend/.goku/generated/service/core/mthd"
	svcuser_mthd "sampleapp/backend/.goku/generated/service/user/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

// NewClient creates a new client for the app, incorportaing direct method clients for each service
func NewClient(ctx context.Context) (app_client.Client, error) {

	var err error
	var c app_client.Client

	// We need to pass the app client for nested methods to be able to call other methods.
	// However, we don't have the app client yet (nothing to pass!). So we will pass a function that will return the app client when called.
	// It is is expected that by the time this function is called, the app client would have been created.
	getClientFn := func(ctx context.Context) app_client.Client {
		log.Debug(ctx, "Fetching app client", "namespaceLevel", "app", "clientType", "method")
		return c
	}

	// Create the method clients for each service
	var req app_clientimpl.NewClientRequest
	{
		svcClient, err := svcauth_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for service [$s]: %w", "Auth", err)
		}
		req.AuthClient = svcClient
	}
	{
		svcClient, err := svccore_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for service [$s]: %w", "Core", err)
		}
		req.CoreClient = svcClient
	}
	{
		svcClient, err := svcuser_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for service [$s]: %w", "User", err)
		}
		req.UserClient = svcClient
	}

	c, err = app_clientimpl.NewClient(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Creating new client: %w", err)
	}

	return c, nil
}
