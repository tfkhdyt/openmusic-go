//go:build wireinject
// +build wireinject

package playlistsongactivity

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/playlistsongactivity"
)

func InitializeService() *Service {
	wire.Build(
		NewService,
		playlistsongactivity.NewRepository,
	)

	return nil
}
