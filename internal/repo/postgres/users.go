package postgres

import "database/sql"

type userDB struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
