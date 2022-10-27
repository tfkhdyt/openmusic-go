//go:build wireinject
// +build wireinject

package playlist

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/playlist"
)

func InitializeService() *Service {
	wire.Build(
		NewService,
		playlist.NewRepository,
	)

	return nil
}
