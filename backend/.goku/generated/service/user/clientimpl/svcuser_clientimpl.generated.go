package svcuser_clientimpl

import (
	"context"

	"github.com/teejays/gokutil/log"

	svcuser_client "sampleapp/backend/.goku/generated/service/user/client"
	svcuser_entteam_client "sampleapp/backend/.goku/generated/service/user/entity/team/client"
	svcuser_entuser_client "sampleapp/backend/.goku/generated/service/user/entity/user/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svcuser")

// client is an implementation of the Client interface. It is simply a wrapper around the child clients,
// which should be implemented separately.
type client struct {
	svcuser_client.EntitiesClient
	svcuser_client.CustomClient
}

// NewClientRequest is the request type for a new client.
type NewClientRequest struct {
	EntitiesClient svcuser_client.EntitiesClient
	CustomClient   svcuser_client.CustomClient
}

func NewClient(ctx context.Context, req NewClientRequest) (svcuser_client.Client, error) {
	cl := client{
		EntitiesClient: req.EntitiesClient,
		CustomClient:   req.CustomClient,
	}
	return cl, nil
}

// Service Entities Client
// entitiesClient is a collection of all entity clients under the service.
type entitiesClient struct {
	teamClient svcuser_entteam_client.Client
	userClient svcuser_entuser_client.Client
}

func (c entitiesClient) Team() svcuser_entteam_client.Client {
	return c.teamClient
}
func (c entitiesClient) User() svcuser_entuser_client.Client {
	return c.userClient
}

// NewEntitiesClientRequest is the request type for a new client.
type NewEntitiesClientRequest struct {
	TeamClient svcuser_entteam_client.Client
	UserClient svcuser_entuser_client.Client
}

func NewEntitiesClient(ctx context.Context, req NewEntitiesClientRequest) (svcuser_client.EntitiesClient, error) {
	cl := entitiesClient{
		teamClient: req.TeamClient,
		userClient: req.UserClient,
	}
	return cl, nil
}

// Custom Client: These include the actual logic implementation -- so they should be implemented by specific client implementations.
