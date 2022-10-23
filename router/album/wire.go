//go:build wireinject
// +build wireinject

package album

import (
	"github.com/google/wire"
	albumsController "github.com/tfkhdyt/openmusic-go/controller/album"
	albumsService "github.com/tfkhdyt/openmusic-go/service/album"
)

func InitializeRouter() *Router {
	wire.Build(
		NewRouter,
		albumsController.NewController,
		albumsService.InitializeService,
	)

	return nil
}
