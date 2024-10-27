package svccore_entsecretdecryptable_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_entsecretdecryptable_client "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/client"
	svccore_entsecretdecryptable_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/clientimpl"
	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svccore_entsecretdecryptable")

// NewClient creates a new instance of the entity client.
func NewClient(ctx context.Context, baseURL string) (svccore_entsecretdecryptable_client.Client, error) {
	httpCrudClient, err := clienttyp.NewEntityHTTPClient[svccore_entsecretdecryptable_typ.SecretDecryptable, svccore_entsecretdecryptable_typ.SecretDecryptableField, svccore_entsecretdecryptable_typ.SecretDecryptableFilter](ctx, baseURL)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	cl, err := svccore_entsecretdecryptable_clientimpl.NewClient(ctx,
		svccore_entsecretdecryptable_clientimpl.NewClientRequest{
			CRUDClient: httpCrudClient,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("creating new entity client: %w", err)
	}

	return cl, nil
}
