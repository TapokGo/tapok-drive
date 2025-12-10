package main

import (
	"log"

	"github.com/TapokGo/tapok-drive/internal/app"
	"github.com/TapokGo/tapok-drive/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	app.Run()
}
