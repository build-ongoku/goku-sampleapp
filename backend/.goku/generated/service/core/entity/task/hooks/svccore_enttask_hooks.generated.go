package svccore_enttask_hooks

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svccore_enttask_meta "sampleapp/backend/.goku/generated/service/core/entity/task/meta"
	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"
)

func InitializeHooks(ctx context.Context, c app_client.Client) error {
	// Hook Save Pre
	{
		err := svccore_enttask_meta.GetEntityDALMeta().TypeDALMeta.SetHookSavePre(
			func(ctx context.Context, v svccore_enttask_typ.Task) (svccore_enttask_typ.Task, error) {
				return c.Core().Task().HookSavePre(ctx, v)
			},
		)
		if err != nil {
			return fmt.Errorf("SettingSavePrehook: %w", err)
		}
	}

	return nil
}
