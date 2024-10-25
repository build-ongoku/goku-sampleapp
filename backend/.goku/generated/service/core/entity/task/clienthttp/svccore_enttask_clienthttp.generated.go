package svccore_enttask_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_enttask_client "sampleapp/backend/.goku/generated/service/core/entity/task/client"
	svccore_enttask_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/task/clientimpl"
	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svccore_enttask")

// NewClient creates a new instance of the entity client.
func NewClient(ctx context.Context, baseURL string) (svccore_enttask_client.Client, error) {
	httpCrudClient, err := clienttyp.NewEntityHTTPClient[svccore_enttask_typ.Task, svccore_enttask_typ.TaskField, svccore_enttask_typ.TaskFilter](ctx, baseURL)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	cl, err := svccore_enttask_clientimpl.NewClient(ctx,
		svccore_enttask_clientimpl.NewClientRequest{
			CRUDClient: httpCrudClient,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("creating new entity client: %w", err)
	}

	return cl, nil
}
