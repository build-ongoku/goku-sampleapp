package svccore_client

import (
	"context"
	"net/http"

	"github.com/teejays/gokutil/log"

	svccore_entfile_client "sampleapp/backend/.goku/generated/service/core/entity/file/client"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
	svccore_entjobapplicant_client "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/client"
	svccore_entsecretdecryptable_client "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/client"
	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
	svccore_entsecretkey_client "sampleapp/backend/.goku/generated/service/core/entity/secret_key/client"
	svccore_enttask_client "sampleapp/backend/.goku/generated/service/core/entity/task/client"
	svccore_enttaskrun_client "sampleapp/backend/.goku/generated/service/core/entity/task_run/client"
	svccore_typ "sampleapp/backend/.goku/generated/service/core/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore")

// Service Client

// Client is the interface for any Core service client.
// A service is made up of it's constitient entity clients, and any service level custom methods.
type Client interface {
	EntitiesClient
	CustomClient
}

// EntitiesClient provides access to the clients of entities in this service.
type EntitiesClient interface {
	File() svccore_entfile_client.Client
	JobApplicant() svccore_entjobapplicant_client.Client
	SecretDecryptable() svccore_entsecretdecryptable_client.Client
	SecretKey() svccore_entsecretkey_client.Client
	Task() svccore_enttask_client.Client
	TaskRun() svccore_enttaskrun_client.Client
}

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
	FileUpload(ctx context.Context, req *http.Request) (svccore_entfile_typ.File, error)
	SecretDecryptabeAdd(ctx context.Context, req svccore_typ.SecretDecryptableAddRequest) (svccore_entsecretdecryptable_typ.SecretDecryptable, error)
}
