package svccore_enttask_client

import (
	"context"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore_enttask")

// Client is the interface that provides access to all the entity's methods.
type Client interface {
	CRUDClient
	CustomClient
}

// CRUDClient interface provides access to all the basic CRUD methods.
type CRUDClient = clienttyp.EntityBaseClient[svccore_enttask_typ.Task, svccore_enttask_typ.TaskInput, svccore_enttask_typ.TaskField, svccore_enttask_typ.TaskFilter]

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
	ActionRun(ctx context.Context, req svccore_enttask_typ.StandardEntityRequest) (svccore_enttask_typ.StandardEntityResponse, error)
	HookSavePre(ctx context.Context, req svccore_enttask_typ.Task) (svccore_enttask_typ.Task, error)
}
