package svccore_client

import (
	"context"
	"net/http"

	"github.com/teejays/gokutil/log"

	svccore_entfile_client "sampleapp/backend/.goku/generated/service/core/entity/file/client"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
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
}

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
	FileUpload(ctx context.Context, req *http.Request) (svccore_entfile_typ.File, error)
}
