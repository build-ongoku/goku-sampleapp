package svcauth_entsecret_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svcauth_entsecret_client "sampleapp/backend/.goku/generated/service/auth/entity/secret/client"
	svcauth_entsecret_clientimpl "sampleapp/backend/.goku/generated/service/auth/entity/secret/clientimpl"
	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svcauth_entsecret")

// NewClient creates a new instance of the entity client.
func NewClient(ctx context.Context, baseURL string) (svcauth_entsecret_client.Client, error) {
	httpCrudClient, err := clienttyp.NewEntityHTTPClient[svcauth_entsecret_typ.Secret, svcauth_entsecret_typ.SecretField, svcauth_entsecret_typ.SecretFilter](ctx, baseURL)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	cl, err := svcauth_entsecret_clientimpl.NewClient(ctx,
		svcauth_entsecret_clientimpl.NewClientRequest{
			CRUDClient: httpCrudClient,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("creating new entity client: %w", err)
	}

	return cl, nil
}
