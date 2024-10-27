package svccore_entsecretdecryptable_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entsecretdecryptable_client "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/client"
	svccore_entsecretdecryptable_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/clientimpl"
	svccore_entsecretdecryptable_dal "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/dal"
	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
	sttc_svccore_entsecretdecryptable_mthd "sampleapp/backend/.goku/static/service/core/secret_decryptable/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svccore_entsecretdecryptable_client.Client, error) {

	crudClient, err := svccore_entsecretdecryptable_dal.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Creating new DAL client: %w", err)
	}

	cl, err := svccore_entsecretdecryptable_clientimpl.NewClient(ctx, svccore_entsecretdecryptable_clientimpl.NewClientRequest{
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

func (c customClient) ActionDecrypt(ctx context.Context, req svccore_entsecretdecryptable_typ.StandardEntityRequest) (svccore_entsecretdecryptable_typ.StandardEntityResponse, error) {

	var resp svccore_entsecretdecryptable_typ.StandardEntityResponse

	var err error

	llog.Info(ctx, "START ActionDecrypt", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	resp, err = sttc_svccore_entsecretdecryptable_mthd.ActionDecrypt(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END ActionDecrypt", "response", resp)

	return resp, err
}

func (c customClient) HookCreatePre(ctx context.Context, req svccore_entsecretdecryptable_typ.SecretDecryptable) (svccore_entsecretdecryptable_typ.SecretDecryptable, error) {

	var resp svccore_entsecretdecryptable_typ.SecretDecryptable

	var err error

	llog.Info(ctx, "START HookCreatePre", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	resp, err = sttc_svccore_entsecretdecryptable_mthd.HookCreatePre(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END HookCreatePre", "response", resp)

	return resp, err
}
