package auth

import (
	"github.com/tfkhdyt/openmusic-go/service/auth"
	"github.com/tfkhdyt/openmusic-go/service/user"
	"github.com/tfkhdyt/openmusic-go/tokenize/token"
)

type Controller struct {
	authService  *auth.Service
	usersService *user.Service
	tokenManager *token.Manager
}

func NewController(authService *auth.Service, usersService *user.Service, tokenManager *token.Manager) *Controller {
	return &Controller{authService, usersService, tokenManager}
}
