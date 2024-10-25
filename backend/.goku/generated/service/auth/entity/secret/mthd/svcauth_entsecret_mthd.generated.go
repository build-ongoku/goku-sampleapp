package svcauth_entsecret_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_entsecret_client "sampleapp/backend/.goku/generated/service/auth/entity/secret/client"
	svcauth_entsecret_clientimpl "sampleapp/backend/.goku/generated/service/auth/entity/secret/clientimpl"
	svcauth_entsecret_dal "sampleapp/backend/.goku/generated/service/auth/entity/secret/dal"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svcauth_entsecret_client.Client, error) {

	crudClient, err := svcauth_entsecret_dal.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Creating new DAL client: %w", err)
	}

	cl, err := svcauth_entsecret_clientimpl.NewClient(ctx, svcauth_entsecret_clientimpl.NewClientRequest{
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
