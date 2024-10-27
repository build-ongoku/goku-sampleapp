package svccore_entsecretkey_init

import (
	"context"
	app_client "sampleapp/backend/.goku/generated/client"
	sttc_svccore_entsecretkey_mthd "sampleapp/backend/.goku/static/service/core/secret_key/mthd"
)

func Init(ctx context.Context, c app_client.Client) error {
	err := sttc_svccore_entsecretkey_mthd.HookInit(ctx, c)
	if err != nil {
		return err
	}

	return nil
}
