//go:build wireinject
// +build wireinject

package album

import (
	"github.com/google/wire"
	postgresConfig "github.com/tfkhdyt/openmusic-go/config/postgres"
	postgresDb "github.com/tfkhdyt/openmusic-go/db/postgres"
	albumRepository "github.com/tfkhdyt/openmusic-go/repository/postgres/album"
	albumService "github.com/tfkhdyt/openmusic-go/service/album"
)

func InitializeController() *Controller {
	wire.Build(
		NewController,
		albumService.NewService,
		albumRepository.NewRepository,
		postgresDb.NewDB,
		postgresConfig.NewConfig,
	)

	return nil
}
