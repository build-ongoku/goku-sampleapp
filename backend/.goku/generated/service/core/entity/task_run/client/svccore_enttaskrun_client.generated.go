package svccore_enttaskrun_client

import (
	"context"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svccore_enttaskrun")

// Client is the interface that provides access to all the entity's methods.
type Client interface {
	CRUDClient
	CustomClient
}

// CRUDClient interface provides access to all the basic CRUD methods.
type CRUDClient = clienttyp.EntityBaseClient[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunInput, svccore_enttaskrun_typ.TaskRunField, svccore_enttaskrun_typ.TaskRunFilter]

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
	ActionRun(ctx context.Context, req svccore_enttaskrun_typ.StandardEntityRequest) (svccore_enttaskrun_typ.StandardEntityResponse, error)
	HookCreatePre(ctx context.Context, req svccore_enttaskrun_typ.TaskRun) (svccore_enttaskrun_typ.TaskRun, error)
}
