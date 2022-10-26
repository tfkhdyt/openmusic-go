//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/auth"
)

func InitializeService() *Service {
	wire.Build(
		NewService,
		auth.NewRepository,
	)

	return nil
}
