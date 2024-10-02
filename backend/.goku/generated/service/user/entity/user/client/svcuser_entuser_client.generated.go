package svcuser_entuser_client

import (
	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svcuser_entuser")

// Client is the interface that provides access to all the entity's methods.
type Client interface {
	CRUDClient
	CustomClient
}

// CRUDClient interface provides access to all the basic CRUD methods.
type CRUDClient = clienttyp.EntityBaseClient[svcuser_entuser_typ.User, svcuser_entuser_typ.UserInput, svcuser_entuser_typ.UserField, svcuser_entuser_typ.UserFilter]

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
}
