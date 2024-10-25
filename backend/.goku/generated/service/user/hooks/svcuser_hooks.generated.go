package svcuser_hooks

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svcuser_entteam_hooks "sampleapp/backend/.goku/generated/service/user/entity/team/hooks"
	svcuser_entuser_hooks "sampleapp/backend/.goku/generated/service/user/entity/user/hooks"
)

func InitializeHooks(ctx context.Context, c app_client.Client) error {
	// Team
	{
		err := svcuser_entteam_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Entity [Team]: %w", err)
		}
	}
	// User
	{
		err := svcuser_entuser_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Entity [User]: %w", err)
		}
	}
	return nil
}
