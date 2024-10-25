package svccore_entfile_client

import (
	"context"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore_entfile")

// Client is the interface that provides access to all the entity's methods.
type Client interface {
	CRUDClient
	CustomClient
}

// CRUDClient interface provides access to all the basic CRUD methods.
type CRUDClient = clienttyp.EntityBaseClient[svccore_entfile_typ.File, svccore_entfile_typ.FileInput, svccore_entfile_typ.FileField, svccore_entfile_typ.FileFilter]

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
	HookCreatePre(ctx context.Context, req svccore_entfile_typ.File) (svccore_entfile_typ.File, error)
}
