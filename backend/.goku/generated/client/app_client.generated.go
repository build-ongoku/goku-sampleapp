package app_client

import (
	"context"

	"github.com/teejays/gokutil/log"

	svcauth_client "sampleapp/backend/.goku/generated/service/auth/client"
	svccore_client "sampleapp/backend/.goku/generated/service/core/client"
	svcuser_client "sampleapp/backend/.goku/generated/service/user/client"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("app")

// ComboClient is the interface for the entire app. It includes all possible clients.
type ComboClient interface {
	Client // default
	Http() Client
	Method() Client
}

// Client is the interface for the entire app.
// This allows a service to connect and communicate with other services.
type Client interface {
	Auth() svcauth_client.Client
	Core() svccore_client.Client
	User() svcuser_client.Client
}

type ClientGetterFn func(ctx context.Context) Client
