package postgres

import (
	"database/sql"
	"fmt"
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
func (u *userRepository) Create(email, passwordHash string) error {
	q := `
	INSERT INTO users(email, password_hash) 
	VALUES ($1, $2)
	ON CONFLICT (email) DO NOTHING
	`

	res, err := u.db.Exec(q, email, passwordHash)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rowa affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrUserDuplicate
	}

	return nil
}
