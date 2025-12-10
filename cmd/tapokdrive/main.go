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
		log.Fatalf("%s", err.Error())
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	app.Run()
}
