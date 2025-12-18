package service

import (
	"context"

	"github.com/google/uuid"
)

// FileRepository temp
type FileRepository interface{}

// FolderRepository
type FolderRepository interface{}

// UserRepository
type UserRepository interface {
	Create(ctx context.Context, id uuid.UUID, email, passwordHah string) error
}
