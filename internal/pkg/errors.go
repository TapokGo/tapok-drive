package pkg

import "errors"

var (
	ErrDuplicate error = errors.New("user with this email already exists")
)
