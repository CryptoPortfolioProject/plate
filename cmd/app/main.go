package main

import (
	"context"
	"gateway/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("app not running: %v", err)
	}
}
