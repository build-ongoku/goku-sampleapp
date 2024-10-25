package svccore_entfile_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entfile_client "sampleapp/backend/.goku/generated/service/core/entity/file/client"
	svccore_entfile_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/file/clientimpl"
	svccore_entfile_dal "sampleapp/backend/.goku/generated/service/core/entity/file/dal"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
	sttc_svccore_entfile_mthd "sampleapp/backend/.goku/static/service/core/file/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svccore_entfile_client.Client, error) {

	crudClient, err := svccore_entfile_dal.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Creating new DAL client: %w", err)
	}

	cl, err := svccore_entfile_clientimpl.NewClient(ctx, svccore_entfile_clientimpl.NewClientRequest{
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

func (c customClient) HookCreatePre(ctx context.Context, req svccore_entfile_typ.File) (svccore_entfile_typ.File, error) {

	var resp svccore_entfile_typ.File

	var err error

	llog.Info(ctx, "START HookCreatePre", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	resp, err = sttc_svccore_entfile_mthd.HookCreatePre(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END HookCreatePre", "response", resp)

	return resp, err
}
