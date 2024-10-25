package svccore_entfile_typhelper

import (
	"context"
	"fmt"
	svccore_entfile_client "sampleapp/backend/.goku/generated/service/core/entity/file/client"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"

	"github.com/teejays/gokutil/scalars"
)

func GetByID(ctx context.Context, c svccore_entfile_client.CRUDClient, id scalars.ID) (svccore_entfile_typ.File, error) {
	var out svccore_entfile_typ.File
	if id.IsEmpty() {
		return out, fmt.Errorf("ID field is empty")
	}
	getReq := svccore_entfile_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
