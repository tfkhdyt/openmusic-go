//go:build wireinject
// +build wireinject

package song

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/song"
)

func InitializeService() *Service {
	wire.Build(
		NewService,
		song.NewRepository,
	)

	return nil
}
