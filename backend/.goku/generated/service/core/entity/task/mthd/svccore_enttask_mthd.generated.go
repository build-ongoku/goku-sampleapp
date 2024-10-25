package svccore_enttask_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_enttask_client "sampleapp/backend/.goku/generated/service/core/entity/task/client"
	svccore_enttask_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/task/clientimpl"
	svccore_enttask_dal "sampleapp/backend/.goku/generated/service/core/entity/task/dal"
	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"
	custom_svccore_enttask_mthd "sampleapp/backend/src/service/core/task/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svccore_enttask_client.Client, error) {

	crudClient, err := svccore_enttask_dal.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Creating new DAL client: %w", err)
	}

	cl, err := svccore_enttask_clientimpl.NewClient(ctx, svccore_enttask_clientimpl.NewClientRequest{
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

func (c customClient) ActionRun(ctx context.Context, req svccore_enttask_typ.StandardEntityRequest) (svccore_enttask_typ.StandardEntityResponse, error) {

	var resp svccore_enttask_typ.StandardEntityResponse

	var err error

	llog.Info(ctx, "START ActionRun", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC
	resp, err = custom_svccore_enttask_mthd.ActionRun(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END ActionRun", "response", resp)

	return resp, err
}

func (c customClient) HookSavePre(ctx context.Context, req svccore_enttask_typ.Task) (svccore_enttask_typ.Task, error) {

	var resp svccore_enttask_typ.Task

	var err error

	llog.Info(ctx, "START HookSavePre", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC
	resp, err = custom_svccore_enttask_mthd.HookSavePre(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END HookSavePre", "response", resp)

	return resp, err
}
