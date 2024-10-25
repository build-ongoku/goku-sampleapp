package svccore_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	svccore_client "sampleapp/backend/.goku/generated/service/core/client"
	svccore_clientimpl "sampleapp/backend/.goku/generated/service/core/clientimpl"
	svccore_entfile_clienthttp "sampleapp/backend/.goku/generated/service/core/entity/file/clienthttp"
	svccore_enttask_clienthttp "sampleapp/backend/.goku/generated/service/core/entity/task/clienthttp"
	svccore_enttaskrun_clienthttp "sampleapp/backend/.goku/generated/service/core/entity/task_run/clienthttp"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svccore")

func NewClient(ctx context.Context, baseURL string) (svccore_client.Client, error) {

	// Create Entities Client
	var entitiesClientReq svccore_clientimpl.NewEntitiesClientRequest
	{
		entityClient, err := svccore_entfile_clienthttp.NewClient(ctx, baseURL)
		if err != nil {
			return nil, fmt.Errorf("Create client for entity [File] : %w", err)
		}
		entitiesClientReq.FileClient = entityClient
	}
	{
		entityClient, err := svccore_enttask_clienthttp.NewClient(ctx, baseURL)
		if err != nil {
			return nil, fmt.Errorf("Create client for entity [Task] : %w", err)
		}
		entitiesClientReq.TaskClient = entityClient
	}
	{
		entityClient, err := svccore_enttaskrun_clienthttp.NewClient(ctx, baseURL)
		if err != nil {
			return nil, fmt.Errorf("Create client for entity [TaskRun] : %w", err)
		}
		entitiesClientReq.TaskRunClient = entityClient
	}
	entitiesClient, err := svccore_clientimpl.NewEntitiesClient(ctx, entitiesClientReq)
	if err != nil {
		return nil, fmt.Errorf("Create entities client: %w", err)
	}

	// Client for the service
	var clientReq svccore_clientimpl.NewClientRequest
	clientReq.EntitiesClient = entitiesClient
	// TODO: clientReq.CustomClient = customClient{...}

	cl, err := svccore_clientimpl.NewClient(ctx, clientReq)
	if err != nil {
		return nil, err
	}

	return cl, err
}
