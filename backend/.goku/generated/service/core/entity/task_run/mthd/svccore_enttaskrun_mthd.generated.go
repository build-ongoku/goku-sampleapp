package svccore_enttaskrun_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_enttaskrun_client "sampleapp/backend/.goku/generated/service/core/entity/task_run/client"
	svccore_enttaskrun_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/task_run/clientimpl"
	svccore_enttaskrun_dal "sampleapp/backend/.goku/generated/service/core/entity/task_run/dal"
	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"
	custom_svccore_enttaskrun_mthd "sampleapp/backend/src/service/core/task_run/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svccore_enttaskrun_client.Client, error) {

	crudClient, err := svccore_enttaskrun_dal.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Creating new DAL client: %w", err)
	}

	cl, err := svccore_enttaskrun_clientimpl.NewClient(ctx, svccore_enttaskrun_clientimpl.NewClientRequest{
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

func (c customClient) ActionRun(ctx context.Context, req svccore_enttaskrun_typ.StandardEntityRequest) (svccore_enttaskrun_typ.StandardEntityResponse, error) {

	var resp svccore_enttaskrun_typ.StandardEntityResponse

	var err error

	llog.Info(ctx, "START ActionRun", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC
	resp, err = custom_svccore_enttaskrun_mthd.ActionRun(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END ActionRun", "response", resp)

	return resp, err
}

func (c customClient) HookCreatePre(ctx context.Context, req svccore_enttaskrun_typ.TaskRun) (svccore_enttaskrun_typ.TaskRun, error) {

	var resp svccore_enttaskrun_typ.TaskRun

	var err error

	llog.Info(ctx, "START HookCreatePre", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC
	resp, err = custom_svccore_enttaskrun_mthd.HookCreatePre(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END HookCreatePre", "response", resp)

	return resp, err
}
