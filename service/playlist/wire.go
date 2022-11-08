//go:build wireinject
// +build wireinject

package playlist

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/playlist"
	"github.com/tfkhdyt/openmusic-go/service/collab"
)

func InitializeService() *Service {
	wire.Build(
		NewService,
		playlist.NewRepository,
		collab.InitializeService,
	)

	return nil
}
