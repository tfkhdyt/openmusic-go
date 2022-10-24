//go:build wireinject
// +build wireinject

package album

import (
	"github.com/google/wire"
	albumsController "github.com/tfkhdyt/openmusic-go/controller/album"
	albumsService "github.com/tfkhdyt/openmusic-go/service/album"
	"github.com/tfkhdyt/openmusic-go/service/song"
)

func InitializeRouter() *Router {
	wire.Build(
		NewRouter,
		albumsController.NewController,
		albumsService.InitializeService,
		song.InitializeService,
	)

	return nil
}
