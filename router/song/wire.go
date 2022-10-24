//go:build wireinject
// +build wireinject

package song

import (
	"github.com/google/wire"
	songsController "github.com/tfkhdyt/openmusic-go/controller/song"
	songsService "github.com/tfkhdyt/openmusic-go/service/song"
)

func InitializeRouter() *Router {
	wire.Build(
		NewRouter,
		songsController.NewController,
		songsService.InitializeService,
	)

	return nil
}
