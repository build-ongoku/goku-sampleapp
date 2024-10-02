package svcuser_entteam_clientimpl

import (
	"context"

	"github.com/teejays/gokutil/log"

	svcuser_entteam_client "sampleapp/backend/.goku/generated/service/user/entity/team/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svcuser_entteam")

// client is an implementation of the entity Client. It is simply a wrapper around the child clients,
// which should be implemented separately.
type client struct {
	svcuser_entteam_client.CRUDClient
	svcuser_entteam_client.CustomClient
}

// NewClientRequest is the request type for a new client.
type NewClientRequest struct {
	CRUDClient   svcuser_entteam_client.CRUDClient
	CustomClient svcuser_entteam_client.CustomClient
}

func NewClient(ctx context.Context, req NewClientRequest) (svcuser_entteam_client.Client, error) {
	cl := client{
		CRUDClient:   req.CRUDClient,
		CustomClient: req.CustomClient,
	}
	return cl, nil
}
