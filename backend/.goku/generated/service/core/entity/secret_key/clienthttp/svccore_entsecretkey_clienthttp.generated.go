package svccore_entsecretkey_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_entsecretkey_client "sampleapp/backend/.goku/generated/service/core/entity/secret_key/client"
	svccore_entsecretkey_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/secret_key/clientimpl"
	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svccore_entsecretkey")

// NewClient creates a new instance of the entity client.
func NewClient(ctx context.Context, baseURL string) (svccore_entsecretkey_client.Client, error) {
	httpCrudClient, err := clienttyp.NewEntityHTTPClient[svccore_entsecretkey_typ.SecretKey, svccore_entsecretkey_typ.SecretKeyField, svccore_entsecretkey_typ.SecretKeyFilter](ctx, baseURL)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	cl, err := svccore_entsecretkey_clientimpl.NewClient(ctx,
		svccore_entsecretkey_clientimpl.NewClientRequest{
			CRUDClient: httpCrudClient,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("creating new entity client: %w", err)
	}

	return cl, nil
}
