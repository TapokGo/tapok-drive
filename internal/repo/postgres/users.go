package postgres

import "database/sql"

type userDB struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

// UserRepository is a model of user repo
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository returns a user db utilities
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
