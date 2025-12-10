// Package service provides ulilities for work with users, files and folders
package service

type userService struct{}

// NewUserService returns user services utilities
func NewUserService(UserRepository) *userService {
	return &userService{}
}
