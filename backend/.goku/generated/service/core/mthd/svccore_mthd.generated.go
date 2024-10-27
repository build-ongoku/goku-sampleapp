package svccore_mthd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_client "sampleapp/backend/.goku/generated/service/core/client"
	svccore_clientimpl "sampleapp/backend/.goku/generated/service/core/clientimpl"
	svccore_entfile_mthd "sampleapp/backend/.goku/generated/service/core/entity/file/mthd"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
	svccore_entjobapplicant_mthd "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/mthd"
	svccore_entsecretdecryptable_mthd "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/mthd"
	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
	svccore_entsecretkey_mthd "sampleapp/backend/.goku/generated/service/core/entity/secret_key/mthd"
	svccore_enttask_mthd "sampleapp/backend/.goku/generated/service/core/entity/task/mthd"
	svccore_enttaskrun_mthd "sampleapp/backend/.goku/generated/service/core/entity/task_run/mthd"
	svccore_typ "sampleapp/backend/.goku/generated/service/core/typ"
	sttc_svccore_mthd "sampleapp/backend/.goku/static/service/core/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svccore_client.Client, error) {

	// Req for creating new svc client
	newEntitiesClientReq := svccore_clientimpl.NewEntitiesClientRequest{}
	{
		entClient, err := svccore_entfile_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for entity [File]: %w", err)
		}
		newEntitiesClientReq.FileClient = entClient
	}
	{
		entClient, err := svccore_entjobapplicant_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for entity [JobApplicant]: %w", err)
		}
		newEntitiesClientReq.JobApplicantClient = entClient
	}
	{
		entClient, err := svccore_entsecretdecryptable_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for entity [SecretDecryptable]: %w", err)
		}
		newEntitiesClientReq.SecretDecryptableClient = entClient
	}
	{
		entClient, err := svccore_entsecretkey_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for entity [SecretKey]: %w", err)
		}
		newEntitiesClientReq.SecretKeyClient = entClient
	}
	{
		entClient, err := svccore_enttask_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for entity [Task]: %w", err)
		}
		newEntitiesClientReq.TaskClient = entClient
	}
	{
		entClient, err := svccore_enttaskrun_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for entity [TaskRun]: %w", err)
		}
		newEntitiesClientReq.TaskRunClient = entClient
	}
	entitiesClient, err := svccore_clientimpl.NewEntitiesClient(ctx, newEntitiesClientReq)
	if err != nil {
		return nil, fmt.Errorf("Creating new entities client: %w", err)
	}

	// Create the full service client
	cl, err := svccore_clientimpl.NewClient(ctx, svccore_clientimpl.NewClientRequest{
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

func (c customClient) FileUpload(ctx context.Context, req *http.Request) (svccore_entfile_typ.File, error) {

	var resp svccore_entfile_typ.File

	var err error

	llog.Info(ctx, "START FileUpload")
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	resp, err = sttc_svccore_mthd.FileUpload(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END FileUpload", "response", resp)

	return resp, err
}

func (c customClient) SecretDecryptabeAdd(ctx context.Context, req svccore_typ.SecretDecryptableAddRequest) (svccore_entsecretdecryptable_typ.SecretDecryptable, error) {

	var resp svccore_entsecretdecryptable_typ.SecretDecryptable

	var err error

	llog.Info(ctx, "START SecretDecryptabeAdd", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	resp, err = sttc_svccore_mthd.SecretDecryptabeAdd(ctx, c, req)
	if err != nil {
		return resp, err
	}

	llog.Info(ctx, "END SecretDecryptabeAdd", "response", resp)

	return resp, err
}
