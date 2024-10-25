package svccore_entfile_hooks

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entfile_meta "sampleapp/backend/.goku/generated/service/core/entity/file/meta"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
)

func InitializeHooks(ctx context.Context, c app_client.Client) error {
	// Hook Create Pre
	{
		err := svccore_entfile_meta.GetEntityDALMeta().TypeDALMeta.SetHookCreatePre(
			func(ctx context.Context, v svccore_entfile_typ.File) (svccore_entfile_typ.File, error) {
				return c.Core().File().HookCreatePre(ctx, v)
			},
		)
		if err != nil {
			return fmt.Errorf("SettingCreatePrehook: %w", err)
		}
	}

	return nil
}
