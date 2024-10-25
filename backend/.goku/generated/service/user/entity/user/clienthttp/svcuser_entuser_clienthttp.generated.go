package svcuser_entuser_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svcuser_entuser_client "sampleapp/backend/.goku/generated/service/user/entity/user/client"
	svcuser_entuser_clientimpl "sampleapp/backend/.goku/generated/service/user/entity/user/clientimpl"
	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svcuser_entuser")

// NewClient creates a new instance of the entity client.
func NewClient(ctx context.Context, baseURL string) (svcuser_entuser_client.Client, error) {
	httpCrudClient, err := clienttyp.NewEntityHTTPClient[svcuser_entuser_typ.User, svcuser_entuser_typ.UserField, svcuser_entuser_typ.UserFilter](ctx, baseURL)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	cl, err := svcuser_entuser_clientimpl.NewClient(ctx,
		svcuser_entuser_clientimpl.NewClientRequest{
			CRUDClient: httpCrudClient,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("creating new entity client: %w", err)
	}

	return cl, nil
}
