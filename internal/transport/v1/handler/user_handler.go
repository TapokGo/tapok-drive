// Package hanlder provides server routes 
package handler

import "github.com/TapokGo/tapok-drive/internal/transport"

type handler struct{}

func NewUserhandler(transport.UserService) *handler {
	return &handler{}
}
