package svcauth_dal

import (
	"context"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // required for 'postgres' dialect
	"github.com/teejays/gokutil/log"

	"github.com/teejays/gokutil/client/db"

	app_dal "sampleapp/backend/.goku/generated/dal"
)

var llog = log.GetLogger().WithHeading("DAL").WithHeading("$.service[Auth]")

/* Type Methods */

// DAL encapsulates DAL methods for types that fall under Auth
type DAL struct {
	app_dal.DAL

	conn db.Connection
}

func NewDAL(ctx context.Context, conn db.Connection) (DAL, error) {
	dal := DAL{conn: conn}
	parentDAL, err := app_dal.NewDAL(ctx, conn)
	if err != nil {
		return DAL{}, err
	}
	dal.DAL = parentDAL
	return dal, nil
}
