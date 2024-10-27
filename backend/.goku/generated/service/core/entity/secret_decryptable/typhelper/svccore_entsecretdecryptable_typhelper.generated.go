package svccore_entsecretdecryptable_typhelper

import (
	"context"
	"fmt"
	svccore_entsecretdecryptable_client "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/client"
	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
	svccore_entsecretkey_client "sampleapp/backend/.goku/generated/service/core/entity/secret_key/client"
	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"

	"github.com/teejays/gokutil/scalars"
)

func GetByID(ctx context.Context, c svccore_entsecretdecryptable_client.CRUDClient, id scalars.ID) (svccore_entsecretdecryptable_typ.SecretDecryptable, error) {
	var out svccore_entsecretdecryptable_typ.SecretDecryptable
	if id.IsEmpty() {
		return out, fmt.Errorf("ID field is empty")
	}
	getReq := svccore_entsecretdecryptable_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
func GetSecretKey(ctx context.Context, c svccore_entsecretkey_client.CRUDClient, ent svccore_entsecretdecryptable_typ.SecretDecryptable) (svccore_entsecretkey_typ.SecretKey, error) {
	var out svccore_entsecretkey_typ.SecretKey
	id := ent.SecretKeyID
	if id.IsEmpty() {
		return out, fmt.Errorf("SecretKeyID field is empty")
	}
	getReq := svccore_entsecretkey_typ.GetRequest{}
	getReq.ID = id
	return c.Get(ctx, getReq)
}
