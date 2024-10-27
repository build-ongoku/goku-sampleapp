package svccore_hooks

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entfile_hooks "sampleapp/backend/.goku/generated/service/core/entity/file/hooks"
	svccore_entjobapplicant_hooks "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/hooks"
	svccore_entsecretdecryptable_hooks "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/hooks"
	svccore_entsecretkey_hooks "sampleapp/backend/.goku/generated/service/core/entity/secret_key/hooks"
	svccore_enttask_hooks "sampleapp/backend/.goku/generated/service/core/entity/task/hooks"
	svccore_enttaskrun_hooks "sampleapp/backend/.goku/generated/service/core/entity/task_run/hooks"
)

func InitializeHooks(ctx context.Context, c app_client.Client) error {
	// File
	{
		err := svccore_entfile_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Entity [File]: %w", err)
		}
	}
	// Job Applicant
	{
		err := svccore_entjobapplicant_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Entity [JobApplicant]: %w", err)
		}
	}
	// Secret Decryptable
	{
		err := svccore_entsecretdecryptable_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Entity [SecretDecryptable]: %w", err)
		}
	}
	// Secret Key
	{
		err := svccore_entsecretkey_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Entity [SecretKey]: %w", err)
		}
	}
	// Task
	{
		err := svccore_enttask_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Entity [Task]: %w", err)
		}
	}
	// Task Run
	{
		err := svccore_enttaskrun_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Entity [TaskRun]: %w", err)
		}
	}
	return nil
}
