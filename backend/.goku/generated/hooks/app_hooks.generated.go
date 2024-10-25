package app_hooks

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_hooks "sampleapp/backend/.goku/generated/service/auth/hooks"
	svccore_hooks "sampleapp/backend/.goku/generated/service/core/hooks"
	svcuser_hooks "sampleapp/backend/.goku/generated/service/user/hooks"
)

func InitializeHooks(ctx context.Context, c app_client.Client) error {
	// Auth
	{
		err := svcauth_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Service [Auth]: %w", err)
		}
	}
	// Core
	{
		err := svccore_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Service [Core]: %w", err)
		}
	}
	// User
	{
		err := svcuser_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Service [User]: %w", err)
		}
	}
	return nil
}
