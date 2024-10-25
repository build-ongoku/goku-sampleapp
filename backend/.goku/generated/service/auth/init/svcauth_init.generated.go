package svcauth_init

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_entsecret_init "sampleapp/backend/.goku/generated/service/auth/entity/secret/init"
)

func Init(ctx context.Context, c app_client.Client) error {
	{
		err := svcauth_entsecret_init.Init(ctx, c)
		if err != nil {
			return fmt.Errorf("Initializing entity [Secret]: %w", err)
		}
	}

	return nil
}
