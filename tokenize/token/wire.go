//go:build wireinject
// +build wireinject

package token

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/config/jwt"
)

func InitializeManager() *Manager {
	wire.Build(
		NewManager,
		jwt.NewConfig,
	)

	return nil
}
