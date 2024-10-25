package svcauth_hooks

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_entsecret_hooks "sampleapp/backend/.goku/generated/service/auth/entity/secret/hooks"
)

func InitializeHooks(ctx context.Context, c app_client.Client) error {
	// Secret
	{
		err := svcauth_entsecret_hooks.InitializeHooks(ctx, c)
		if err != nil {
			return fmt.Errorf("Entity [Secret]: %w", err)
		}
	}
	return nil
}
