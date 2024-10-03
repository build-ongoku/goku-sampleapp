package svccore_hooks

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entfile_hooks "sampleapp/backend/.goku/generated/service/core/entity/file/hooks"
	svccore_entjobapplicant_hooks "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/hooks"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svccore")

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
	return nil
}
