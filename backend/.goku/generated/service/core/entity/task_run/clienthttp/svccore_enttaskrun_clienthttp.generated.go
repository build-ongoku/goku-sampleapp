package svccore_enttaskrun_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svccore_enttaskrun_client "sampleapp/backend/.goku/generated/service/core/entity/task_run/client"
	svccore_enttaskrun_clientimpl "sampleapp/backend/.goku/generated/service/core/entity/task_run/clientimpl"
	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svccore_enttaskrun")

// NewClient creates a new instance of the entity client.
func NewClient(ctx context.Context, baseURL string) (svccore_enttaskrun_client.Client, error) {
	httpCrudClient, err := clienttyp.NewEntityHTTPClient[svccore_enttaskrun_typ.TaskRun, svccore_enttaskrun_typ.TaskRunField, svccore_enttaskrun_typ.TaskRunFilter](ctx, baseURL)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	cl, err := svccore_enttaskrun_clientimpl.NewClient(ctx,
		svccore_enttaskrun_clientimpl.NewClientRequest{
			CRUDClient: httpCrudClient,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("creating new entity client: %w", err)
	}

	return cl, nil
}
