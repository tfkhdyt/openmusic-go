//go:build wireinject
// +build wireinject

package collab

import (
	"github.com/google/wire"
	collabsController "github.com/tfkhdyt/openmusic-go/controller/collab"
	colllabsService "github.com/tfkhdyt/openmusic-go/service/collab"
	"github.com/tfkhdyt/openmusic-go/service/playlist"
	"github.com/tfkhdyt/openmusic-go/service/user"
)

func InitializeRouter() *Router {
	wire.Build(
		NewRouter,
		collabsController.NewController,
		colllabsService.InitializeService,
		playlist.InitializeService,
		user.InitializeService,
	)

	return nil
}
