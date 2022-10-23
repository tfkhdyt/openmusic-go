// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package postgres

import (
	"github.com/tfkhdyt/openmusic-go/config/postgres"
)

// Injectors from wire.go:

func InitializeDB() *DB {
	config := postgres.NewConfig()
	db := NewDB(config)
	return db
}