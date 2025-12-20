// Package transport provides interfaces of service layer
package transport

import (
	"context"

	"github.com/TapokGo/tapok-drive/internal/service"
)

// UserService temp
type UserService interface {
	Create(ctx context.Context, dto service.CreateUser) (*service.UserResponse, error)
}
