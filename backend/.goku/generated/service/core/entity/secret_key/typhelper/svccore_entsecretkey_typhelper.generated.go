package svccore_entsecretkey_typhelper

import (
	"context"
	"fmt"
	svccore_entsecretkey_client "sampleapp/backend/.goku/generated/service/core/entity/secret_key/client"
	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"

	"github.com/teejays/gokutil/scalars"
)

func GetByID(ctx context.Context, c svccore_entsecretkey_client.CRUDClient, id scalars.ID) (svccore_entsecretkey_typ.SecretKey, error) {
	var out svccore_entsecretkey_typ.SecretKey
	if id.IsEmpty() {
		return out, fmt.Errorf("ID field is empty")
	}
	getReq := svccore_entsecretkey_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
