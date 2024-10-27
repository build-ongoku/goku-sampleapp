package svccore_entsecretkey_client

import (
	"context"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore_entsecretkey")

// Client is the interface that provides access to all the entity's methods.
type Client interface {
	CRUDClient
	CustomClient
}

// CRUDClient interface provides access to all the basic CRUD methods.
type CRUDClient = clienttyp.EntityBaseClient[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyInput, svccore_entsecretkey_typ.SecretKeyField, svccore_entsecretkey_typ.SecretKeyFilter]

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
	HookCreatePre(ctx context.Context, req svccore_entsecretkey_typ.SecretKey) (svccore_entsecretkey_typ.SecretKey, error)
	HookInit(ctx context.Context) error
}
