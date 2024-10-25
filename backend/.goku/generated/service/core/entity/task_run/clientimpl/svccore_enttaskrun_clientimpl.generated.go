package svccore_enttaskrun_clientimpl

import (
	"context"

	"github.com/teejays/gokutil/log"

	svccore_enttaskrun_client "sampleapp/backend/.goku/generated/service/core/entity/task_run/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore_enttaskrun")

// client is an implementation of the entity Client. It is simply a wrapper around the child clients,
// which should be implemented separately.
type client struct {
	svccore_enttaskrun_client.CRUDClient
	svccore_enttaskrun_client.CustomClient
}

// NewClientRequest is the request type for a new client.
type NewClientRequest struct {
	CRUDClient   svccore_enttaskrun_client.CRUDClient
	CustomClient svccore_enttaskrun_client.CustomClient
}

func NewClient(ctx context.Context, req NewClientRequest) (svccore_enttaskrun_client.Client, error) {
	cl := client{
		CRUDClient:   req.CRUDClient,
		CustomClient: req.CustomClient,
	}
	return cl, nil
}
