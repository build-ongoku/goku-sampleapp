package app_init

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_init "sampleapp/backend/.goku/generated/service/auth/init"
	svccore_init "sampleapp/backend/.goku/generated/service/core/init"
	svcuser_init "sampleapp/backend/.goku/generated/service/user/init"
)

func Init(ctx context.Context, c app_client.Client) error {
	{
		err := svcauth_init.Init(ctx, c)
		if err != nil {
			return fmt.Errorf("Initializing svc [Auth]: %w", err)
		}
	}
	{
		err := svccore_init.Init(ctx, c)
		if err != nil {
			return fmt.Errorf("Initializing svc [Core]: %w", err)
		}
	}
	{
		err := svcuser_init.Init(ctx, c)
		if err != nil {
			return fmt.Errorf("Initializing svc [User]: %w", err)
		}
	}

	return nil
}
