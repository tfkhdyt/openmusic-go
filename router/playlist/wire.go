//go:build wireinject
// +build wireinject

package playlist

import (
	"github.com/google/wire"
	playlistsController "github.com/tfkhdyt/openmusic-go/controller/playlist"
	playlistSongsController "github.com/tfkhdyt/openmusic-go/controller/playlistsong"
	playlistsService "github.com/tfkhdyt/openmusic-go/service/playlist"
	playlistSongsService "github.com/tfkhdyt/openmusic-go/service/playlistsong"
	"github.com/tfkhdyt/openmusic-go/service/song"
)

func InitializeRouter() *Router {
	wire.Build(
		NewRouter,
		playlistsController.NewController,
		playlistsService.InitializeService,
		playlistSongsController.NewController,
		playlistSongsService.InitializeService,
		song.InitializeService,
	)

	return nil
}
