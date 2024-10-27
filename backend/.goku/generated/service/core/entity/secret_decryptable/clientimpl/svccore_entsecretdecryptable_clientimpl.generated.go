package svccore_entsecretdecryptable_clientimpl

import (
	"context"

	"github.com/teejays/gokutil/log"

	svccore_entsecretdecryptable_client "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore_entsecretdecryptable")

// client is an implementation of the entity Client. It is simply a wrapper around the child clients,
// which should be implemented separately.
type client struct {
	svccore_entsecretdecryptable_client.CRUDClient
	svccore_entsecretdecryptable_client.CustomClient
}

// NewClientRequest is the request type for a new client.
type NewClientRequest struct {
	CRUDClient   svccore_entsecretdecryptable_client.CRUDClient
	CustomClient svccore_entsecretdecryptable_client.CustomClient
}

func NewClient(ctx context.Context, req NewClientRequest) (svccore_entsecretdecryptable_client.Client, error) {
	cl := client{
		CRUDClient:   req.CRUDClient,
		CustomClient: req.CustomClient,
	}
	return cl, nil
}
