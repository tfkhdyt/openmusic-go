//go:build wireinject
// +build wireinject

package postgres

import (
	"github.com/google/wire"
	"github.com/tfkhdyt/openmusic-go/config/postgres"
)

func InitializeDB() *DB {
	wire.Build(
		NewDB,
		postgres.NewConfig,
	)

	return nil
}
