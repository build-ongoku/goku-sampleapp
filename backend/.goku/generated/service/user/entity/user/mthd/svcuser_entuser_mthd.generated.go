package svcuser_entuser_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svcuser_entuser_client "sampleapp/backend/.goku/generated/service/user/entity/user/client"
	svcuser_entuser_clientimpl "sampleapp/backend/.goku/generated/service/user/entity/user/clientimpl"
	svcuser_entuser_dal "sampleapp/backend/.goku/generated/service/user/entity/user/dal"
	sttc_svcuser_entuser_mthd "sampleapp/backend/.goku/static/service/user/user/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svcuser_entuser_client.Client, error) {

	crudClient, err := svcuser_entuser_dal.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Creating new DAL client: %w", err)
	}

	cl, err := svcuser_entuser_clientimpl.NewClient(ctx, svcuser_entuser_clientimpl.NewClientRequest{
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

func (c customClient) HookInit(ctx context.Context) error {

	var err error

	llog.Info(ctx, "START HookInit")
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	err = sttc_svcuser_entuser_mthd.HookInit(ctx, c)
	if err != nil {
		return err
	}

	llog.Info(ctx, "END HookInit")

	return err
}
