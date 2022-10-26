//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
	usersController "github.com/tfkhdyt/openmusic-go/controller/user"
	usersService "github.com/tfkhdyt/openmusic-go/service/user"
)

func InitializeRouter() *Router {
	wire.Build(
		NewRouter,
		usersController.NewController,
		usersService.InitializeService,
	)

	return nil
}
