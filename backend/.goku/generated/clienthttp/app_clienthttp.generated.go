package app_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	app_clientimpl "sampleapp/backend/.goku/generated/clientimpl"
	svcauth_clienthttp "sampleapp/backend/.goku/generated/service/auth/clienthttp"
	svccore_clienthttp "sampleapp/backend/.goku/generated/service/core/clienthttp"
	svcuser_clienthttp "sampleapp/backend/.goku/generated/service/user/clienthttp"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("app")

// NewClient creates a new instance of the http based app client.
func NewClient(ctx context.Context, baseURL string) (app_client.Client, error) {
	newClientReq := app_clientimpl.NewClientRequest{}
	{
		svcClient, err := svcauth_clienthttp.NewClient(ctx, baseURL)
		if err != nil {
			return nil, fmt.Errorf("Creating HTTP client for service [%s]: %w", "Auth", err)
		}
		newClientReq.AuthClient = svcClient
	}
	{
		svcClient, err := svccore_clienthttp.NewClient(ctx, baseURL)
		if err != nil {
			return nil, fmt.Errorf("Creating HTTP client for service [%s]: %w", "Core", err)
		}
		newClientReq.CoreClient = svcClient
	}
	{
		svcClient, err := svcuser_clienthttp.NewClient(ctx, baseURL)
		if err != nil {
			return nil, fmt.Errorf("Creating HTTP client for service [%s]: %w", "User", err)
		}
		newClientReq.UserClient = svcClient
	}
	cl, err := app_clientimpl.NewClient(ctx, newClientReq)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	return cl, nil
}
