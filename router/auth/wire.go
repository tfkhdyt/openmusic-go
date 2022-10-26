//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/google/wire"
	authController "github.com/tfkhdyt/openmusic-go/controller/auth"
	authService "github.com/tfkhdyt/openmusic-go/service/auth"
	"github.com/tfkhdyt/openmusic-go/service/user"
	"github.com/tfkhdyt/openmusic-go/tokenize/token"
)

func InitializeRouter() *Router {
	wire.Build(
		NewRouter,
		authController.NewController,
		authService.InitializeService,
		user.InitializeService,
		token.InitializeManager,
	)

	return nil
}
