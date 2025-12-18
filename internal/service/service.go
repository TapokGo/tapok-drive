package service

import "errors"

var (
	ErrShortPassword error = errors.New("password must be at leats 8")
	ErrWeakPassword  error = errors.New("password is weak")
	ErrInvalidEmail  error = errors.New("invalid email")
)

type userService struct {
	repo UserRepository
}

// NewUserService returns user services utilities
func NewUserService(repo UserRepository) *userService {
	return &userService{repo: repo}
}
