package mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/env"
	"github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/log"
	"github.com/teejays/gokutil/scalars"

	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_typ "sampleapp/backend/.goku/generated/service/auth/typ"
	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"
	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"
	app_typ "sampleapp/backend/.goku/generated/typ"
)

var (
	DEV_USER_EMAIL    scalars.Email = scalars.NewEmail("dev@goku.com")
	DEV_USER_PASSWORD string        = "password"
)

func HookInit(ctx context.Context, c app_client.Client) error {

	if env.GetEnv() == env.DEV || env.GetEnv() == env.TEST {
		err := UpsertDevUser(ctx, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpsertDevUser(ctx context.Context, c app_client.Client) error {

	// Get any existing user
	resp, err := c.User().User().List(ctx, svcuser_entuser_typ.ListRequest{
		Filter: svcuser_entuser_typ.UserFilter{
			Email: filter.NewEmailCondition(filter.EQUAL, DEV_USER_EMAIL),
		},
	})
	if err != nil {
		return fmt.Errorf("Fetching User with email [%s]: %w", DEV_USER_EMAIL, err)
	}
	if resp.Count > 0 {
		// User already exists
		log.Info(ctx, "Dev User already exists.")
		return nil
	}

	// Create an team
	team, err := c.User().Team().Add(ctx, svcuser_entteam_typ.AddRequest{
		Object: svcuser_entteam_typ.TeamInput{
			Name: "Goku",
		},
	})
	if err != nil {
		return fmt.Errorf("Creating Team: %w", err)
	}

	// Register the user
	_, err = c.Auth().RegisterUser(ctx, svcauth_typ.RegisterUserRequest{
		Email:    DEV_USER_EMAIL,
		Password: DEV_USER_PASSWORD,
		Name: app_typ.PersonNameInput{
			FirstName: "Dev",
			LastName:  "User",
		},
		TeamID: team.ID,
	})
	if err != nil {
		return fmt.Errorf("Registering Dev User: %w", err)
	}

	return nil
}
