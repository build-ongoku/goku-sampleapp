package app_clientimpl

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_client "sampleapp/backend/.goku/generated/service/auth/client"
	svccore_client "sampleapp/backend/.goku/generated/service/core/client"
	svcuser_client "sampleapp/backend/.goku/generated/service/user/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("app")

// client is an implementation of the Client interface.
type client struct {
	authClient svcauth_client.Client
	coreClient svccore_client.Client
	userClient svcuser_client.Client
}

func (c client) Auth() svcauth_client.Client {
	return c.authClient
}
func (c client) Core() svccore_client.Client {
	return c.coreClient
}
func (c client) User() svcuser_client.Client {
	return c.userClient
}

// NewClientRequest is the request type for a new client.
type NewClientRequest struct {
	AuthClient svcauth_client.Client
	CoreClient svccore_client.Client
	UserClient svcuser_client.Client
}

func NewClient(ctx context.Context, req NewClientRequest) (app_client.Client, error) {
	// Ensure that no client is nil
	if req.AuthClient == nil {
		return nil, fmt.Errorf("Client for service [Auth] is nil")
	}
	if req.CoreClient == nil {
		return nil, fmt.Errorf("Client for service [Core] is nil")
	}
	if req.UserClient == nil {
		return nil, fmt.Errorf("Client for service [User] is nil")
	}
	return &client{
		authClient: req.AuthClient,
		coreClient: req.CoreClient,
		userClient: req.UserClient,
	}, nil
}
