// Package postgres provides utilities for work with postgres db
package postgres

type repository struct{}

// New creates new postgres connection
func New() (*repository, error) {
	return &repository{}, nil
}
