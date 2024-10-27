package svccore_entjobapplicant_client

import (
	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_entjobapplicant_typ "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore_entjobapplicant")

// Client is the interface that provides access to all the entity's methods.
type Client interface {
	CRUDClient
	CustomClient
}

// CRUDClient interface provides access to all the basic CRUD methods.
type CRUDClient = clienttyp.EntityBaseClient[svccore_entjobapplicant_typ.JobApplicant, svccore_entjobapplicant_typ.JobApplicantInput, svccore_entjobapplicant_typ.JobApplicantField, svccore_entjobapplicant_typ.JobApplicantFilter]

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
}
