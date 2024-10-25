package svcuser_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	svcuser_client "sampleapp/backend/.goku/generated/service/user/client"
	svcuser_clientimpl "sampleapp/backend/.goku/generated/service/user/clientimpl"
	svcuser_entteam_clienthttp "sampleapp/backend/.goku/generated/service/user/entity/team/clienthttp"
	svcuser_entuser_clienthttp "sampleapp/backend/.goku/generated/service/user/entity/user/clienthttp"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svcuser")

func NewClient(ctx context.Context, baseURL string) (svcuser_client.Client, error) {

	// Create Entities Client
	var entitiesClientReq svcuser_clientimpl.NewEntitiesClientRequest
	{
		entityClient, err := svcuser_entteam_clienthttp.NewClient(ctx, baseURL)
		if err != nil {
			return nil, fmt.Errorf("Create client for entity [Team] : %w", err)
		}
		entitiesClientReq.TeamClient = entityClient
	}
	{
		entityClient, err := svcuser_entuser_clienthttp.NewClient(ctx, baseURL)
		if err != nil {
			return nil, fmt.Errorf("Create client for entity [User] : %w", err)
		}
		entitiesClientReq.UserClient = entityClient
	}
	entitiesClient, err := svcuser_clientimpl.NewEntitiesClient(ctx, entitiesClientReq)
	if err != nil {
		return nil, fmt.Errorf("Create entities client: %w", err)
	}

	// Client for the service
	var clientReq svcuser_clientimpl.NewClientRequest
	clientReq.EntitiesClient = entitiesClient
	// TODO: clientReq.CustomClient = customClient{...}

	cl, err := svcuser_clientimpl.NewClient(ctx, clientReq)
	if err != nil {
		return nil, err
	}

	return cl, err
}
