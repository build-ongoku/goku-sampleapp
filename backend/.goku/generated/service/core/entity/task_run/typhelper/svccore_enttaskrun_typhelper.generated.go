package svccore_enttaskrun_typhelper

import (
	"context"
	"fmt"
	svccore_enttask_client "sampleapp/backend/.goku/generated/service/core/entity/task/client"
	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"
	svccore_enttaskrun_client "sampleapp/backend/.goku/generated/service/core/entity/task_run/client"
	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"

	"github.com/teejays/gokutil/scalars"
)

func GetByID(ctx context.Context, c svccore_enttaskrun_client.CRUDClient, id scalars.ID) (svccore_enttaskrun_typ.TaskRun, error) {
	var out svccore_enttaskrun_typ.TaskRun
	if id.IsEmpty() {
		return out, fmt.Errorf("ID field is empty")
	}
	getReq := svccore_enttaskrun_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
func GetTask(ctx context.Context, c svccore_enttask_client.CRUDClient, ent svccore_enttaskrun_typ.TaskRun) (svccore_enttask_typ.Task, error) {
	var out svccore_enttask_typ.Task
	id := ent.TaskID
	if id.IsEmpty() {
		return out, fmt.Errorf("TaskID field is empty")
	}
	getReq := svccore_enttask_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
