package main

import (
	"context"
	"os"

	"github.com/teejays/gokutil/log"

	"sampleapp/backend/.goku/static/monoservice"
)

func main() {
	ctx := context.Background()
	if err := monoservice.Run(ctx); err != nil {
		log.Error(ctx, "Command failed", "error", err)
		os.Exit(1)
	}
}
