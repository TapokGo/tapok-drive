// Package main is main package of tapok-drive app
package main

import (
	"log"

	"github.com/TapokGo/tapok-drive/internal/app"
	"github.com/TapokGo/tapok-drive/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := app.Close(); err != nil {
			app.Logger.Error("failed to close dependencies", "error", err)
		}
	}()

	if err := app.Run(); err != nil {
		app.Logger.Error("application failed", "error", err)
	}
}
