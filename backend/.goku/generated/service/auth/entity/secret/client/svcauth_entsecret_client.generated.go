package svcauth_entsecret_client

import (
	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svcauth_entsecret")

// Client is the interface that provides access to all the entity's methods.
type Client interface {
	CRUDClient
	CustomClient
}

// CRUDClient interface provides access to all the basic CRUD methods.
type CRUDClient = clienttyp.EntityBaseClient[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretInput, svcauth_entsecret_typ.SecretField, svcauth_entsecret_typ.SecretFilter]

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
}
