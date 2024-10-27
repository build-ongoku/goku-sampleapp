package svccore_entjobapplicant_typhelper

import (
	"context"
	"fmt"
	svccore_entjobapplicant_client "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/client"
	svccore_entjobapplicant_typ "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/typ"

	"github.com/teejays/gokutil/scalars"
)

func GetByID(ctx context.Context, c svccore_entjobapplicant_client.CRUDClient, id scalars.ID) (svccore_entjobapplicant_typ.JobApplicant, error) {
	var out svccore_entjobapplicant_typ.JobApplicant
	if id.IsEmpty() {
		return out, fmt.Errorf("ID field is empty")
	}
	getReq := svccore_entjobapplicant_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
