package svccore_entjobapplicant_clientimpl

import (
	"context"

	"github.com/teejays/gokutil/log"

	svccore_entjobapplicant_client "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore_entjobapplicant")

// client is an implementation of the entity Client. It is simply a wrapper around the child clients,
// which should be implemented separately.
type client struct {
	svccore_entjobapplicant_client.CRUDClient
	svccore_entjobapplicant_client.CustomClient
}

// NewClientRequest is the request type for a new client.
type NewClientRequest struct {
	CRUDClient   svccore_entjobapplicant_client.CRUDClient
	CustomClient svccore_entjobapplicant_client.CustomClient
}

func NewClient(ctx context.Context, req NewClientRequest) (svccore_entjobapplicant_client.Client, error) {
	cl := client{
		CRUDClient:   req.CRUDClient,
		CustomClient: req.CustomClient,
	}
	return cl, nil
}
