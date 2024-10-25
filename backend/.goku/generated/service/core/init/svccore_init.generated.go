package svccore_init

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entfile_init "sampleapp/backend/.goku/generated/service/core/entity/file/init"
	svccore_enttask_init "sampleapp/backend/.goku/generated/service/core/entity/task/init"
	svccore_enttaskrun_init "sampleapp/backend/.goku/generated/service/core/entity/task_run/init"
)

func Init(ctx context.Context, c app_client.Client) error {
	{
		err := svccore_entfile_init.Init(ctx, c)
		if err != nil {
			return fmt.Errorf("Initializing entity [File]: %w", err)
		}
	}
	{
		err := svccore_enttask_init.Init(ctx, c)
		if err != nil {
			return fmt.Errorf("Initializing entity [Task]: %w", err)
		}
	}
	{
		err := svccore_enttaskrun_init.Init(ctx, c)
		if err != nil {
			return fmt.Errorf("Initializing entity [TaskRun]: %w", err)
		}
	}

	return nil
}
