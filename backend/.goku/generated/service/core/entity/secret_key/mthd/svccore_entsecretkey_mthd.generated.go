package svccore_entsecretkey_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entsecretkey_client "sampleapp/backend/.goku/generated/service/core/entity/secret_key/client"
	svccore_entsecretkey_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/secret_key/clientimpl"
	svccore_entsecretkey_dal "sampleapp/backend/.goku/generated/service/core/entity/secret_key/dal"
	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"
	sttc_svccore_entsecretkey_mthd "sampleapp/backend/.goku/static/service/core/secret_key/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svccore_entsecretkey_client.Client, error) {

	crudClient, err := svccore_entsecretkey_dal.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Creating new DAL client: %w", err)
	}

	cl, err := svccore_entsecretkey_clientimpl.NewClient(ctx, svccore_entsecretkey_clientimpl.NewClientRequest{
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

func (c customClient) HookCreatePre(ctx context.Context, req svccore_entsecretkey_typ.SecretKey) (svccore_entsecretkey_typ.SecretKey, error) {

	var resp svccore_entsecretkey_typ.SecretKey

	var err error

	llog.Info(ctx, "START HookCreatePre", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	resp, err = sttc_svccore_entsecretkey_mthd.HookCreatePre(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END HookCreatePre", "response", resp)

	return resp, err
}

func (c customClient) HookInit(ctx context.Context) error {

	var err error

	llog.Info(ctx, "START HookInit")
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	err = sttc_svccore_entsecretkey_mthd.HookInit(ctx, c)
	if err != nil {
		return err
	}

	llog.Info(ctx, "END HookInit")

	return err
}
