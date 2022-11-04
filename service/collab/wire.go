//go:build wireinject
// +build wireinject

package collab

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/collab"
)

func InitializeService() *Service {
	wire.Build(
		NewService,
		collab.NewRepository,
	)

	return nil
}
