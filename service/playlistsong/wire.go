//go:build wireinject
// +build wireinject

package playlistsong

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/playlistsong"
)

func InitializeService() *Service {
	wire.Build(
		NewService,
		playlistsong.NewRepository,
	)

	return nil
}
