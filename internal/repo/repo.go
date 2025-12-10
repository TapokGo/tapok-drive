// Package repo provides an interface for work with the repository
package repo

type UserRepository interface {
	Create(email, password string) error
}
