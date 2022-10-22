//go:build wireinject
// +build wireinject

package album

import (
	"github.com/google/wire"
	albumRepository "github.com/tfkhdyt/openmusic-go/repository/postgres/album"
	albumService "github.com/tfkhdyt/openmusic-go/service/album"
)

func InitializeController() *Controller {
	wire.Build(
		NewController,
		albumService.NewService,
		albumRepository.NewRepository,
	)

	return nil
}
