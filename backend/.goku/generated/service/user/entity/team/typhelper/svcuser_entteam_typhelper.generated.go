package svcuser_entteam_typhelper

import (
	"context"
	"fmt"
	svcuser_entteam_client "sampleapp/backend/.goku/generated/service/user/entity/team/client"
	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"

	"github.com/teejays/gokutil/scalars"
)

func GetByID(ctx context.Context, c svcuser_entteam_client.CRUDClient, id scalars.ID) (svcuser_entteam_typ.Team, error) {
	var out svcuser_entteam_typ.Team
	if id.IsEmpty() {
		return out, fmt.Errorf("ID field is empty")
	}
	getReq := svcuser_entteam_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
