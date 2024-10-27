package svccore_entsecretdecryptable_client

import (
	"context"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore_entsecretdecryptable")

// Client is the interface that provides access to all the entity's methods.
type Client interface {
	CRUDClient
	CustomClient
}

// CRUDClient interface provides access to all the basic CRUD methods.
type CRUDClient = clienttyp.EntityBaseClient[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableInput, svccore_entsecretdecryptable_typ.SecretDecryptableField, svccore_entsecretdecryptable_typ.SecretDecryptableFilter]

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
	ActionDecrypt(ctx context.Context, req svccore_entsecretdecryptable_typ.StandardEntityRequest) (svccore_entsecretdecryptable_typ.StandardEntityResponse, error)
	HookCreatePre(ctx context.Context, req svccore_entsecretdecryptable_typ.SecretDecryptable) (svccore_entsecretdecryptable_typ.SecretDecryptable, error)
}
