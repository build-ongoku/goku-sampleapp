package svcauth_entsecret_typhelper

import (
	"context"
	"fmt"
	svcauth_entsecret_client "sampleapp/backend/.goku/generated/service/auth/entity/secret/client"
	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"

	"github.com/teejays/gokutil/scalars"
)

func GetByID(ctx context.Context, c svcauth_entsecret_client.CRUDClient, id scalars.ID) (svcauth_entsecret_typ.Secret, error) {
	var out svcauth_entsecret_typ.Secret
	if id.IsEmpty() {
		return out, fmt.Errorf("ID field is empty")
	}
	getReq := svcauth_entsecret_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
