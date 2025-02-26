package mthd

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svc_typ "sampleapp/backend/.goku/generated/service/pluralizer/typ"

	"github.com/teejays/gokutil/strcase"
)

func Pluralize(ctx context.Context, c app_client.Client, req string) (svc_typ.PluralizeResponse, error) {

	if req == "" {
		return svc_typ.PluralizeResponse{}, fmt.Errorf("Empty request")
	}
	p, err := strcase.Pluralize(req)
	return svc_typ.PluralizeResponse{
		Plural: p,
	}, err
}
