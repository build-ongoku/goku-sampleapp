package svccore_entjobapplicant_hooks

import (
	"context"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svccore_entjobapplicant")

func InitializeHooks(ctx context.Context, c app_client.Client) error {

	return nil
}
