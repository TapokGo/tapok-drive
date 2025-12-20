// Package dto provides a response-request models
package dto

import "github.com/google/uuid"

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required, min=8"`
}

type RegisterResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}
