// Package handler provides server routes
package handler

import (
	"encoding/json"
	"io/fs"
	"net/http"

	"github.com/TapokGo/tapok-drive/internal/transport"
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
type Handler struct {
	urlService transport.UserService
	swagger    *Swagger
}

// New return new Handler{}
func New(service transport.UserService, swagger *Swagger) *Handler {
	return &Handler{
		urlService: service,
		swagger:    swagger,
	}
}

// Register regists all routes
func (h *Handler) Register(r chi.Router) {
	r.Use(middleware.Recoverer)
	r.Get("/healthz", h.checkHealth)

	// Swagger
	if h.swagger != nil {
		r.Get("/tapok-drive", h.getSwagger)
		r.Handle("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.FS(h.swagger.SwaggerUI))))
	}
}

// CheckHealth checks server dependencies and return OK
func (h *Handler) checkHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, r, http.StatusOK, "OK")
}

func (h *Handler) getSwagger(w http.ResponseWriter, r *http.Request) {
	logger := middle.FromContext(r.Context())
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

// func writeError(w http.ResponseWriter, err httperror.HTTPError) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(err.Code)
// 	_ = json.NewEncoder(w).Encode(err.Message)
// }
