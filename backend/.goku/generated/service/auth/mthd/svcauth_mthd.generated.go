package svcauth_mthd

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_client "sampleapp/backend/.goku/generated/service/auth/client"
	svcauth_clientimpl "sampleapp/backend/.goku/generated/service/auth/clientimpl"
	svcauth_entsecret_mthd "sampleapp/backend/.goku/generated/service/auth/entity/secret/mthd"
	svcauth_typ "sampleapp/backend/.goku/generated/service/auth/typ"
	sttc_svcauth_mthd "sampleapp/backend/.goku/static/service/auth/mthd"
)

var llog = log.GetLogger().WithHeading("Methods")

func NewClient(ctx context.Context, getClientFn app_client.ClientGetterFn) (svcauth_client.Client, error) {

	// Req for creating new svc client
	newEntitiesClientReq := svcauth_clientimpl.NewEntitiesClientRequest{}
	{
		entClient, err := svcauth_entsecret_mthd.NewClient(ctx, getClientFn)
		if err != nil {
			return nil, fmt.Errorf("Creating new method client for entity [Secret]: %w", err)
		}
		newEntitiesClientReq.SecretClient = entClient
	}
	entitiesClient, err := svcauth_clientimpl.NewEntitiesClient(ctx, newEntitiesClientReq)
	if err != nil {
		return nil, fmt.Errorf("Creating new entities client: %w", err)
	}

	// Create the full service client
	cl, err := svcauth_clientimpl.NewClient(ctx, svcauth_clientimpl.NewClientRequest{
		EntitiesClient: entitiesClient,
		CustomClient:   customClient{getClientFn: getClientFn},
	},
	)
	if err != nil {
		return nil, fmt.Errorf("Creating new client: %w", err)
	}
	return cl, nil
}

// customClient is an implementation of the service's custom methods
type customClient struct {
	getClientFn app_client.ClientGetterFn
	app_client.Client
}

func (c customClient) LoginUser(ctx context.Context, req svcauth_typ.LoginRequest) (svcauth_typ.AuthenticateResponse, error) {
	var resp svcauth_typ.AuthenticateResponse
	var err error

	llog.Info(ctx, "LoginUser invoked", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	resp, err = sttc_svcauth_mthd.LoginUser(ctx, c, req)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func (c customClient) RegisterUser(ctx context.Context, req svcauth_typ.RegisterUserRequest) (svcauth_typ.AuthenticateResponse, error) {
	var resp svcauth_typ.AuthenticateResponse
	var err error

	llog.Info(ctx, "RegisterUser invoked", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	resp, err = sttc_svcauth_mthd.RegisterUser(ctx, c, req)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func (c customClient) AuthenticateToken(ctx context.Context, req svcauth_typ.AuthenticateTokenRequest) (svcauth_typ.AuthenticateResponse, error) {
	var resp svcauth_typ.AuthenticateResponse
	var err error

	llog.Info(ctx, "AuthenticateToken invoked", "request", req)
	// Fetch the app client and make it available under the custom client
	otherC := c.getClientFn(ctx)
	c.Client = otherC

	resp, err = sttc_svcauth_mthd.AuthenticateToken(ctx, c, req)
	if err != nil {
		return resp, err
	}

	return resp, err
}
