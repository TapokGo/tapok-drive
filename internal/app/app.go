// Package app provides utilities for init and start server
package app

import (
	"fmt"

	"github.com/TapokGo/tapok-drive/internal/config"
	"github.com/TapokGo/tapok-drive/internal/logger"
	"github.com/TapokGo/tapok-drive/internal/repo/postgres"
	"github.com/TapokGo/tapok-drive/internal/service"
	"github.com/TapokGo/tapok-drive/internal/transport/v1/handler"
)

type app struct {
	Logger logger.Logger
	cfg    config.Config
	server Userhandler
}

// New inits app dependencies
func New(cfg config.Config) (*app, error) {
	logger, err := logger.NewSlog("")
	if err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}

	repo, err := postgres.New()
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	service := service.NewUserService(repo)
	handler := handler.NewUserhandler(service)

	app := &app{
		Logger: logger,
		cfg:    cfg,
		server: handler,
	}

	return app, nil
}

// Run starts server
func (a *app) Run() {
	a.Logger.Info(a.cfg.AppEnv)
}
