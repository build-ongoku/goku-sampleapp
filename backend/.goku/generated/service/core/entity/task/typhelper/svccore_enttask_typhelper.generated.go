package svccore_enttask_typhelper

import (
	"context"
	"fmt"
	svccore_enttask_client "sampleapp/backend/.goku/generated/service/core/entity/task/client"
	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"

	"github.com/teejays/gokutil/scalars"
)

func GetByID(ctx context.Context, c svccore_enttask_client.CRUDClient, id scalars.ID) (svccore_enttask_typ.Task, error) {
	var out svccore_enttask_typ.Task
	if id.IsEmpty() {
		return out, fmt.Errorf("ID field is empty")
	}
	getReq := svccore_enttask_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
