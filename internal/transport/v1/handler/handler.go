// Package handler provides server routes
package handler

import (
	"encoding/json"
	"io/fs"
	"net/http"

	"github.com/TapokGo/tapok-drive/internal/transport"
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
func (h *Handler) checkHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, "OK")
}

func (h *Handler) getSwagger(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/yaml")
	w.Write(h.swagger.OpenAPISpec)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// func writeError(w http.ResponseWriter, err httperror.HTTPError) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(err.Code)
// 	_ = json.NewEncoder(w).Encode(err.Message)
// }
