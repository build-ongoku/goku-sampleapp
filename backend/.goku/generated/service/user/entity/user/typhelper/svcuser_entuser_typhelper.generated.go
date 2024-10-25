package svcuser_entuser_typhelper

import (
	"context"
	"fmt"
	svcuser_entuser_client "sampleapp/backend/.goku/generated/service/user/entity/user/client"
	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"

	"github.com/teejays/gokutil/scalars"
)

func GetByID(ctx context.Context, c svcuser_entuser_client.CRUDClient, id scalars.ID) (svcuser_entuser_typ.User, error) {
	var out svcuser_entuser_typ.User
	if id.IsEmpty() {
		return out, fmt.Errorf("ID field is empty")
	}
	getReq := svcuser_entuser_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
