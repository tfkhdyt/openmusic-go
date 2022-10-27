//go:build wireinject
// +build wireinject

package playlist

import (
	"github.com/google/wire"
	playlistsController "github.com/tfkhdyt/openmusic-go/controller/playlist"
	playlistsService "github.com/tfkhdyt/openmusic-go/service/playlist"
)

func InitializeRouter() *Router {
	wire.Build(
		NewRouter,
		playlistsController.NewController,
		playlistsService.InitializeService,
	)

	return nil
}
