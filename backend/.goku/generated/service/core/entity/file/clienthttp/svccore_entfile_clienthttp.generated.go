package svccore_entfile_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_entfile_client "sampleapp/backend/.goku/generated/service/core/entity/file/client"
	svccore_entfile_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/file/clientimpl"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svccore_entfile")

// NewClient creates a new instance of the entity client.
func NewClient(ctx context.Context, baseURL string) (svccore_entfile_client.Client, error) {
	httpCrudClient, err := clienttyp.NewEntityHTTPClient[svccore_entfile_typ.File, svccore_entfile_typ.FileField, svccore_entfile_typ.FileFilter](ctx, baseURL)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	cl, err := svccore_entfile_clientimpl.NewClient(ctx,
		svccore_entfile_clientimpl.NewClientRequest{
			CRUDClient: httpCrudClient,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("creating new entity client: %w", err)
	}

	return cl, nil
}
