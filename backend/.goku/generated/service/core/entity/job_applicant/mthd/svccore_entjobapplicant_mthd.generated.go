package svccore_entjobapplicant_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entjobapplicant_client "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/client"
	svccore_entjobapplicant_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/clientimpl"
	svccore_entjobapplicant_dal "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/dal"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svccore_entjobapplicant_client.Client, error) {

	crudClient, err := svccore_entjobapplicant_dal.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Creating new DAL client: %w", err)
	}

	cl, err := svccore_entjobapplicant_clientimpl.NewClient(ctx, svccore_entjobapplicant_clientimpl.NewClientRequest{
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
