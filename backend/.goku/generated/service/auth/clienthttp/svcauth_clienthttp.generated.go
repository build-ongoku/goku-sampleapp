package svcauth_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	svcauth_client "sampleapp/backend/.goku/generated/service/auth/client"
	svcauth_clientimpl "sampleapp/backend/.goku/generated/service/auth/clientimpl"
	svcauth_entsecret_clienthttp "sampleapp/backend/.goku/generated/service/auth/entity/secret/clienthttp"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svcauth")

func NewClient(ctx context.Context, baseURL string) (svcauth_client.Client, error) {

	// Create Entities Client
	var entitiesClientReq svcauth_clientimpl.NewEntitiesClientRequest
	{
		entityClient, err := svcauth_entsecret_clienthttp.NewClient(ctx, baseURL)
		if err != nil {
			return nil, fmt.Errorf("Create client for entity [Secret] : %w", err)
		}
		entitiesClientReq.SecretClient = entityClient
	}
	entitiesClient, err := svcauth_clientimpl.NewEntitiesClient(ctx, entitiesClientReq)
	if err != nil {
		return nil, fmt.Errorf("Create entities client: %w", err)
	}

	// Client for the service
	var clientReq svcauth_clientimpl.NewClientRequest
	clientReq.EntitiesClient = entitiesClient
	// TODO: clientReq.CustomClient = customClient{...}

	cl, err := svcauth_clientimpl.NewClient(ctx, clientReq)
	if err != nil {
		return nil, err
	}

	return cl, err
}
