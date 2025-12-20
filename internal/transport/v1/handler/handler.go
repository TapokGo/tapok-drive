// Package handler provides server routes
package handler

import (
	"encoding/json"
	"io/fs"
	"net/http"

	"github.com/TapokGo/tapok-drive/internal/logger"
	"github.com/TapokGo/tapok-drive/internal/transport"
	"github.com/TapokGo/tapok-drive/internal/transport/httperror"
	middle "github.com/TapokGo/tapok-drive/internal/transport/v1/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Swagger is a model of swagger data
type Swagger struct {
	OpenAPISpec []byte
	SwaggerUI   fs.FS
}

// Handler is a model of handlers
type handler struct {
	userService transport.UserService
	swagger     *Swagger
	baseLogger  logger.Logger
}

// New return new Handler{}
func New(userService transport.UserService, baseLogger logger.Logger, swagger *Swagger) *handler {
	return &handler{
		userService: userService,
		swagger:     swagger,
		baseLogger:  baseLogger,
	}
}

// Register regists all routes
func (h *handler) Register(r chi.Router) {
	r.Use(middleware.Recoverer)
	r.Use(middle.RequestLogger(h.baseLogger))

	r.Get("/healthz", h.checkHealth)

	// Swagger
	if h.swagger != nil {
		r.Get("/tapok-drive", h.getSwagger)
		r.Handle("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.FS(h.swagger.SwaggerUI))))
	}

	// User routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.registerUser)
	})
}

// CheckHealth checks server dependencies and return OK
func (h *handler) checkHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, r, http.StatusOK, "OK")
}

func (h *handler) getSwagger(w http.ResponseWriter, r *http.Request) {
	logger := middle.FromContext(r.Context())
	logger.Info("open swagger")
	w.Header().Set("Content-Type", "application/yaml")
	_, err := w.Write(h.swagger.OpenAPISpec)
	if err != nil {
		logger.Error("failed to write answer", "error", err)
	}
}

func writeJSON(w http.ResponseWriter, r *http.Request, status int, v any) {
	logger := middle.FromContext(r.Context())
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		logger.Error("failed to encode message", "message", v, "error", err)
	}
}

func writeError(w http.ResponseWriter, r *http.Request, httpError httperror.HTTPError) {
	logger := middle.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpError.Code)
	err := json.NewEncoder(w).Encode(httpError)
	if err != nil {
		logger.Error("failed to encode error message", "error", err)
	}
}
