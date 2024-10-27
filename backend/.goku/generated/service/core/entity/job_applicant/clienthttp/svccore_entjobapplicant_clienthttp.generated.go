package svccore_entjobapplicant_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_entjobapplicant_client "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/client"
	svccore_entjobapplicant_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/clientimpl"
	svccore_entjobapplicant_typ "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/typ"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svccore_entjobapplicant")

// NewClient creates a new instance of the entity client.
func NewClient(ctx context.Context, baseURL string) (svccore_entjobapplicant_client.Client, error) {
	httpCrudClient, err := clienttyp.NewEntityHTTPClient[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantField, svccore_entjobapplicant_typ.JobApplicantFilter](ctx, baseURL)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	cl, err := svccore_entjobapplicant_clientimpl.NewClient(ctx,
		svccore_entjobapplicant_clientimpl.NewClientRequest{
			CRUDClient: httpCrudClient,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("creating new entity client: %w", err)
	}

	return cl, nil
}
