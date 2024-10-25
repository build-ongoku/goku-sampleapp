package svcuser_entuser_init

import (
	"context"
	app_client "sampleapp/backend/.goku/generated/client"
	sttc_svcuser_entuser_mthd "sampleapp/backend/.goku/static/service/user/user/mthd"
)

func Init(ctx context.Context, c app_client.Client) error {
	err := sttc_svcuser_entuser_mthd.HookInit(ctx, c)
	if err != nil {
		return err
	}

	return nil
}
