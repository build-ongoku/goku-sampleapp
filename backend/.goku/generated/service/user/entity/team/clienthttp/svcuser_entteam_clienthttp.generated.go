package svcuser_entteam_clienthttp

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/types/clienttyp"

	svcuser_entteam_client "sampleapp/backend/.goku/generated/service/user/entity/team/client"
	svcuser_entteam_clientimpl "sampleapp/backend/.goku/generated/service/user/entity/team/clientimpl"
	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"
)

var llog = log.GetLogger().WithHeading("Client HTTP").WithHeading("svcuser_entteam")

// NewClient creates a new instance of the entity client.
func NewClient(ctx context.Context, baseURL string) (svcuser_entteam_client.Client, error) {
	httpCrudClient, err := clienttyp.NewEntityHTTPClient[svcuser_entteam_typ.Team, svcuser_entteam_typ.TeamField, svcuser_entteam_typ.TeamFilter](ctx, baseURL)
	if err != nil {
		return nil, fmt.Errorf("creating new entity HTTP client: %w", err)
	}

	cl, err := svcuser_entteam_clientimpl.NewClient(ctx,
		svcuser_entteam_clientimpl.NewClientRequest{
			CRUDClient: httpCrudClient,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("creating new entity client: %w", err)
	}

	return cl, nil
}
