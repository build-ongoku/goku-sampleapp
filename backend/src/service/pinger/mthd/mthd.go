package mthd

import (
	"context"
	app_client "sampleapp/backend/.goku/generated/client"
	svcpinger_typ "sampleapp/backend/.goku/generated/service/pinger/typ"
)

func Ping(ctx context.Context, c app_client.Client, req string) (svcpinger_typ.PingResponse, error) {

	return svcpinger_typ.PingResponse{
		Resp: "Pong",
	}, nil
}
