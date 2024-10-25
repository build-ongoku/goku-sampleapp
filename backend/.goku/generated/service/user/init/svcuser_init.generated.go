package svcuser_init

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svcuser_entteam_init "sampleapp/backend/.goku/generated/service/user/entity/team/init"
	svcuser_entuser_init "sampleapp/backend/.goku/generated/service/user/entity/user/init"
)

func Init(ctx context.Context, c app_client.Client) error {
	{
		err := svcuser_entteam_init.Init(ctx, c)
		if err != nil {
			return fmt.Errorf("Initializing entity [Team]: %w", err)
		}
	}
	{
		err := svcuser_entuser_init.Init(ctx, c)
		if err != nil {
			return fmt.Errorf("Initializing entity [User]: %w", err)
		}
	}

	return nil
}
