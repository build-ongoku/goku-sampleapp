package svcuser_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svcuser_client "sampleapp/backend/.goku/generated/service/user/client"
	svcuser_clientimpl "sampleapp/backend/.goku/generated/service/user/clientimpl"
	svcuser_entteam_mthd "sampleapp/backend/.goku/generated/service/user/entity/team/mthd"
	svcuser_entuser_mthd "sampleapp/backend/.goku/generated/service/user/entity/user/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svcuser_client.Client, error) {

	// Req for creating new svc client
	newEntitiesClientReq := svcuser_clientimpl.NewEntitiesClientRequest{}
	{
		entClient, err := svcuser_entteam_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for entity [Team]: %w", err)
		}
		newEntitiesClientReq.TeamClient = entClient
	}
	{
		entClient, err := svcuser_entuser_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for entity [User]: %w", err)
		}
		newEntitiesClientReq.UserClient = entClient
	}
	entitiesClient, err := svcuser_clientimpl.NewEntitiesClient(ctx, newEntitiesClientReq)
	if err != nil {
		return nil, fmt.Errorf("Creating new entities client: %w", err)
	}

	// Create the full service client
	cl, err := svcuser_clientimpl.NewClient(ctx, svcuser_clientimpl.NewClientRequest{
		EntitiesClient: entitiesClient,
		CustomClient:   customClient{getClientFn: getClientFn},
	},
	)
	if err != nil {
		return nil, fmt.Errorf("Creating new client: %w", err)
	}
	return cl, nil
}

// customClient is an implementation of the service's custom methods
type customClient struct {
	getClientFn app_client.ClientGetterFn
	app_client.Client
}
