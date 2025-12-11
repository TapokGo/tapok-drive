// Package hanlder provides server routes
package handler

import (
	"encoding/json"
	"io/fs"
	"net/http"

	"github.com/TapokGo/tapok-drive/internal/transport"
	"github.com/TapokGo/tapok-drive/internal/transport/httperror"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Swagger is a model of swagger data
type Swagger struct {
	OpenApiSpec []byte
	SwaggerUI   fs.FS
}

// Handler is a model of handlers
type Handler struct {
	urlService transport.UserService
	swagger    *Swagger
}

// NewUserHandler return new Handler{}
func New(service transport.UserService, swagger *Swagger) *Handler {
	return &Handler{
		urlService: service,
		swagger:    swagger,
	}
}

// Register regists all routes
func (h *Handler) Register(r chi.Router) {
	r.Use(middleware.Recoverer)
	r.Get("/healthz", h.CheckHealth)

	// Swagger
	if h.swagger != nil {
		r.Get("/tapok-drive", h.GetSwagger)
		r.Handle("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.FS(h.swagger.SwaggerUI))))
	}
}

// CheckHealth checks server dependencies and return OK
func (h *Handler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, "OK")
}

func (h *Handler) GetSwagger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/yaml")
	w.Write(h.swagger.OpenApiSpec)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, err httperror.HTTPError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	_ = json.NewEncoder(w).Encode(err.Message)
}
