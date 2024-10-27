package svccore_entsecretkey_hooks

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entsecretkey_meta "sampleapp/backend/.goku/generated/service/core/entity/secret_key/meta"
	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"
)

func InitializeHooks(ctx context.Context, c app_client.Client) error {
	// Hook Create Pre
	{
		err := svccore_entsecretkey_meta.GetEntityDALMeta().TypeDALMeta.SetHookCreatePre(
			func(ctx context.Context, v svccore_entsecretkey_typ.SecretKey) (svccore_entsecretkey_typ.SecretKey, error) {
				return c.Core().SecretKey().HookCreatePre(ctx, v)
			},
		)
		if err != nil {
			return fmt.Errorf("SettingCreatePrehook: %w", err)
		}
	}

	return nil
}
