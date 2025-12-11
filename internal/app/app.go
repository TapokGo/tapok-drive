// Package app provides utilities for init and start server
package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/TapokGo/tapok-drive/internal/config"
	"github.com/TapokGo/tapok-drive/internal/logger"
	"github.com/TapokGo/tapok-drive/internal/repo/postgres"
	"github.com/TapokGo/tapok-drive/internal/service"
	"github.com/TapokGo/tapok-drive/internal/transport/v1/handler"
	"github.com/go-chi/chi/v5"
)

type app struct {
	Logger logger.Logger
	cfg    config.Config
	router http.Handler
}

// New inits app dependencies
func New(cfg config.Config) (*app, error) {
	// Init logger
	logger, err := logger.NewSlog(cfg.LogPath, cfg.AppEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}

	// Init repository
	repo, err := postgres.New()
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	// Init service
	userService := service.NewUserService(repo)

	// Init router
	handlers := handler.New(userService)
	r := chi.NewRouter()
	handlers.Register(r)

	app := &app{
		Logger: logger,
		cfg:    cfg,
		router: r,
	}

	return app, nil
}

// Run starts server
func (a *app) Run() error {
	addr := a.cfg.ServerAddress + ":" + strconv.Itoa(a.cfg.ServerPort)
	server := &http.Server{
		Addr:              addr,
		Handler:           a.router,
		ReadTimeout:       a.cfg.Timeout,
		WriteTimeout:      a.cfg.Timeout,
		ReadHeaderTimeout: a.cfg.IdleTimeout,
	}

	a.Logger.Info("Server started", "address", addr)

	// Graceful shutdowm
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	// If get error from server, return server error
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server failed: %w", err)
		}

		return nil

	// If get signal from user of OS use graceful shutdown
	case <-sigCh:
		a.Logger.Info("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			a.Logger.Error("Server shutdown error", "error", err)
			return err
		}

		a.Logger.Info("Server stopped gracefully")
		return nil
	}
}

// Close closes app dependencies
func (a *app) Close() error {
	closeErrors := make([]error, 0, 2)
	// Close log file
	if a.Logger != nil {
		if closer, ok := a.Logger.(io.Closer); ok {
			if err := closer.Close(); err != nil {
				closeErrors = append(closeErrors, err)
			}
		}
		a.Logger = nil
	}

	return errors.Join(closeErrors...)
}
