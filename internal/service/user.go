// Package service provides utilities for work with users, files and folders
package service

import (
	"context"
	"fmt"
	"net/mail"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Email    string
	Password string
}

type UserResponse struct {
	ID    string
	Email string
}

func (u *userService) Create(ctx context.Context, dto CreateUser) (*UserResponse, error) {
	id := uuid.New()
	if err := passwordValidate(dto.Password); err != nil {
		return nil, err
	}

	if err := emailValidate(dto.Email); err != nil {
		return nil, err
	}

	hash, err := hashPassword(dto.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	err = u.repo.Create(ctx, id, dto.Email, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &UserResponse{
		ID:    id.String(),
		Email: dto.Email,
	}, nil
}

func passwordValidate(password string) error {
	if len(password) < 8 {
		return ErrShortPassword
	}

	var hasUpper, hasLower, hasDigital bool
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigital = true
		}
	}

	if !hasUpper || !hasLower || !hasDigital {
		return ErrWeakPassword
	}

	return nil
}

func emailValidate(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
