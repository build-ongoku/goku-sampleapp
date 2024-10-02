package svcauth_clientimpl

import (
	"context"

	"github.com/teejays/gokutil/log"

	svcauth_client "sampleapp/backend/.goku/generated/service/auth/client"
	svcauth_entsecret_client "sampleapp/backend/.goku/generated/service/auth/entity/secret/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svcauth")

// client is an implementation of the Client interface. It is simply a wrapper around the child clients,
// which should be implemented separately.
type client struct {
	svcauth_client.EntitiesClient
	svcauth_client.CustomClient
}

// NewClientRequest is the request type for a new client.
type NewClientRequest struct {
	EntitiesClient svcauth_client.EntitiesClient
	CustomClient   svcauth_client.CustomClient
}

func NewClient(ctx context.Context, req NewClientRequest) (svcauth_client.Client, error) {
	cl := client{
		EntitiesClient: req.EntitiesClient,
		CustomClient:   req.CustomClient,
	}
	return cl, nil
}

// Service Entities Client
// entitiesClient is a collection of all entity clients under the service.
type entitiesClient struct {
	secretClient svcauth_entsecret_client.Client
}

func (c entitiesClient) Secret() svcauth_entsecret_client.Client {
	return c.secretClient
}

// NewEntitiesClientRequest is the request type for a new client.
type NewEntitiesClientRequest struct {
	SecretClient svcauth_entsecret_client.Client
}

func NewEntitiesClient(ctx context.Context, req NewEntitiesClientRequest) (svcauth_client.EntitiesClient, error) {
	cl := entitiesClient{
		secretClient: req.SecretClient,
	}
	return cl, nil
}

// Custom Client: These include the actual logic implementation -- so they should be implemented by specific client implementations.
