package svcuser_entteam_client

import (
	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svcuser_entteam")

// Client is the interface that provides access to all the entity's methods.
type Client interface {
	CRUDClient
	CustomClient
}

// CRUDClient interface provides access to all the basic CRUD methods.
type CRUDClient = clienttyp.EntityBaseClient[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamInput, svcuser_entteam_typ.TeamField, svcuser_entteam_typ.TeamFilter]

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
}
