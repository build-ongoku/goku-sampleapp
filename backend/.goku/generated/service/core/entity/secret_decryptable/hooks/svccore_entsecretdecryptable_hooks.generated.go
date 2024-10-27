package svccore_entsecretdecryptable_hooks

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entsecretdecryptable_meta "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/meta"
	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
)

func InitializeHooks(ctx context.Context, c app_client.Client) error {
	// Hook Create Pre
	{
		err := svccore_entsecretdecryptable_meta.GetEntityDALMeta().TypeDALMeta.SetHookCreatePre(
			func(ctx context.Context, v svccore_entsecretdecryptable_typ.SecretDecryptable) (svccore_entsecretdecryptable_typ.SecretDecryptable, error) {
				return c.Core().SecretDecryptable().HookCreatePre(ctx, v)
			},
		)
		if err != nil {
			return fmt.Errorf("SettingCreatePrehook: %w", err)
		}
	}

	return nil
}
