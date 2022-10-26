//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/user"
)

func InitializeService() *Service {
	wire.Build(
		NewService,
		user.NewRepository,
	)

	return nil
}
