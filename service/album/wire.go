//go:build wireinject
// +build wireinject

package album

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/album"
)

func InitializeService() *Service {
	wire.Build(
		NewService,
		album.NewRepository,
	)

	return nil
}
