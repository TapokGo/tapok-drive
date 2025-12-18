package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// UserRepository is a model of user repo
type userRepository struct {
	db *sql.DB
}

// NewUserRepository returns a user db utilities
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

// Create creates new user
func (u *userRepository) Create(ctx context.Context, id uuid.UUID, email, passwordHash string) error {
	q := `
	INSERT INTO users(id, email, password_hash) 
	VALUES ($1, $2, $3)
	ON CONFLICT (email) DO NOTHING
	`

	res, err := u.db.ExecContext(ctx, q, id, email, passwordHash)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrUserDuplicate
	}

	return nil
}
