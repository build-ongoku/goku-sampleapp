package svcuser_entteam_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svcuser_entteam_client "sampleapp/backend/.goku/generated/service/user/entity/team/client"
	svcuser_entteam_clientimpl "sampleapp/backend/.goku/generated/service/user/entity/team/clientimpl"
	svcuser_entteam_dal "sampleapp/backend/.goku/generated/service/user/entity/team/dal"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svcuser_entteam_client.Client, error) {

	crudClient, err := svcuser_entteam_dal.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Creating new DAL client: %w", err)
	}

	cl, err := svcuser_entteam_clientimpl.NewClient(ctx, svcuser_entteam_clientimpl.NewClientRequest{
		CRUDClient:   crudClient,
		CustomClient: customClient{getClientFn: getClientFn},
	},
	)
	if err != nil {
		return nil, fmt.Errorf("Creating new client: %w", err)
	}
	return cl, nil
}

// customClient is an implementation of the entity's custom method client
type customClient struct {
	getClientFn app_client.ClientGetterFn
	app_client.Client
}
