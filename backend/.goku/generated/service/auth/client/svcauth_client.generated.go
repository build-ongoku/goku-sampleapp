package svcauth_client

import (
	"context"

	"github.com/teejays/gokutil/log"

	svcauth_entsecret_client "sampleapp/backend/.goku/generated/service/auth/entity/secret/client"
	svcauth_typ "sampleapp/backend/.goku/generated/service/auth/typ"
)

var llog = log.GetLogger().WithHeading("Client").WithHeading("svcauth")

// Service Client

// Client is the interface for any Auth service client.
// A service is made up of it's constitient entity clients, and any service level custom methods.
type Client interface {
	EntitiesClient
	CustomClient
}

// EntitiesClient provides access to the clients of entities in this service.
type EntitiesClient interface {
	Secret() svcauth_entsecret_client.Client
}

// CustomClient provides access to all custom methods for this service.
type CustomClient interface {
	AuthenticateToken(ctx context.Context, req svcauth_typ.AuthenticateTokenRequest) (svcauth_typ.AuthenticateResponse, error)
	LoginUser(ctx context.Context, req svcauth_typ.LoginRequest) (svcauth_typ.AuthenticateResponse, error)
	RegisterUser(ctx context.Context, req svcauth_typ.RegisterUserRequest) (svcauth_typ.AuthenticateResponse, error)
}
