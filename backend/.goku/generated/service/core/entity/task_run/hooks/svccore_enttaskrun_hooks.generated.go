package svccore_enttaskrun_hooks

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svccore_enttaskrun_meta "sampleapp/backend/.goku/generated/service/core/entity/task_run/meta"
	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"
)

func InitializeHooks(ctx context.Context, c app_client.Client) error {
	// Hook Create Pre
	{
		err := svccore_enttaskrun_meta.GetEntityDALMeta().TypeDALMeta.SetHookCreatePre(
			func(ctx context.Context, v svccore_enttaskrun_typ.TaskRun) (svccore_enttaskrun_typ.TaskRun, error) {
				return c.Core().TaskRun().HookCreatePre(ctx, v)
			},
		)
		if err != nil {
			return fmt.Errorf("SettingCreatePrehook: %w", err)
		}
	}

	return nil
}
