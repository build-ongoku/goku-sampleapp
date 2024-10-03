package svcuser_entteam_hooks

import (
	"context"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svcuser_entteam")

func InitializeHooks(ctx context.Context, c app_client.Client) error {

	return nil
}
