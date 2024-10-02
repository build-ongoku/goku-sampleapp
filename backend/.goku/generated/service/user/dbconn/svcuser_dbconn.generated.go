package svcuser_dbconn

import (
	"context"
	"fmt"

	"github.com/teejays/gokutil/client/db"
)

// InitDatabaseConnection initializes DB connection for this service.
func InitDatabaseConnection(ctx context.Context, req db.ServiceInitConnectionRequest) error {

	o := db.Options{
		Database: "user",
		Host:     req.Host,
		Port:     req.Port,
		User:     req.User,
		Password: req.Password,
		SSLMode:  "disable",
	}

	err := db.InitDatabase(ctx, o)
	if err != nil {
		return fmt.Errorf("Initializing database connection: %w", err)
	}

	return nil
}
