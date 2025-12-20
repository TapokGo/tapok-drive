package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TapokGo/tapok-drive/internal/service"
	"github.com/TapokGo/tapok-drive/internal/transport/httperror"
	"github.com/TapokGo/tapok-drive/internal/transport/v1/dto"
	middle "github.com/TapokGo/tapok-drive/internal/transport/v1/middleware"
)

func (h *handler) registerUser(w http.ResponseWriter, r *http.Request) {
	logger := middle.FromContext(r.Context())

	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, httperror.InvalidRequest("invalid data"))
		logger.Debug("failed to decode register request body", "error", err)
		return
	}

	logger.Info("received register request")

	serviceDTO := service.CreateUser{
		Email:    req.Email,
		Password: req.Password,
	}

	serviceResp, err := h.userService.Create(r.Context(), serviceDTO)
	if err != nil {
		logger.Error("failed to create user", "error", err)
		if err == service.ErrWeakPassword || err == service.ErrShortPassword {
			writeError(w, r, httperror.InvalidRequest("invalid password"))
			return
		}

		if err == service.ErrInvalidEmail {
			writeError(w, r, httperror.InvalidRequest("invalid email"))
			return
		}

		if err == service.ErrUserExists {
			writeError(w, r, httperror.ConflictError("user with this email already exists"))
			return
		}

		writeError(w, r, httperror.InternalError("failed to register"))
		return
	}

	response := dto.RegisterResponse{
		ID:    serviceResp.ID,
		Email: serviceResp.Email,
	}

	writeJSON(w, r, http.StatusCreated, response)
}
