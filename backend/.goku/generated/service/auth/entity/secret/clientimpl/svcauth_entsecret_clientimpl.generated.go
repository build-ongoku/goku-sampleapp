package svcauth_entsecret_clientimpl

import (
	"context"

	"github.com/teejays/gokutil/log"

	svcauth_entsecret_client "sampleapp/backend/.goku/generated/service/auth/entity/secret/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svcauth_entsecret")

// client is an implementation of the entity Client. It is simply a wrapper around the child clients,
// which should be implemented separately.
type client struct {
	svcauth_entsecret_client.CRUDClient
	svcauth_entsecret_client.CustomClient
}

// NewClientRequest is the request type for a new client.
type NewClientRequest struct {
	CRUDClient   svcauth_entsecret_client.CRUDClient
	CustomClient svcauth_entsecret_client.CustomClient
}

func NewClient(ctx context.Context, req NewClientRequest) (svcauth_entsecret_client.Client, error) {
	cl := client{
		CRUDClient:   req.CRUDClient,
		CustomClient: req.CustomClient,
	}
	return cl, nil
}
