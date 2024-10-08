package svccore_clientimpl

import (
	"context"

	"github.com/teejays/gokutil/log"

	svccore_client "sampleapp/backend/.goku/generated/service/core/client"
	svccore_entfile_client "sampleapp/backend/.goku/generated/service/core/entity/file/client"
	svccore_entjobapplicant_client "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore")

// client is an implementation of the Client interface. It is simply a wrapper around the child clients,
// which should be implemented separately.
type client struct {
	svccore_client.EntitiesClient
	svccore_client.CustomClient
}

// NewClientRequest is the request type for a new client.
type NewClientRequest struct {
	EntitiesClient svccore_client.EntitiesClient
	CustomClient   svccore_client.CustomClient
}

func NewClient(ctx context.Context, req NewClientRequest) (svccore_client.Client, error) {
	cl := client{
		EntitiesClient: req.EntitiesClient,
		CustomClient:   req.CustomClient,
	}
	return cl, nil
}

// Service Entities Client
// entitiesClient is a collection of all entity clients under the service.
type entitiesClient struct {
	fileClient         svccore_entfile_client.Client
	jobApplicantClient svccore_entjobapplicant_client.Client
}

func (c entitiesClient) File() svccore_entfile_client.Client {
	return c.fileClient
}
func (c entitiesClient) JobApplicant() svccore_entjobapplicant_client.Client {
	return c.jobApplicantClient
}

// NewEntitiesClientRequest is the request type for a new client.
type NewEntitiesClientRequest struct {
	FileClient         svccore_entfile_client.Client
	JobApplicantClient svccore_entjobapplicant_client.Client
}

func NewEntitiesClient(ctx context.Context, req NewEntitiesClientRequest) (svccore_client.EntitiesClient, error) {
	cl := entitiesClient{
		fileClient:         req.FileClient,
		jobApplicantClient: req.JobApplicantClient,
	}
	return cl, nil
}

// Custom Client: These include the actual logic implementation -- so they should be implemented by specific client implementations.
