// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package album

import (
	"github.com/tfkhdyt/openmusic-go/config/postgres"
	postgres2 "github.com/tfkhdyt/openmusic-go/db/postgres"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/album"
	album2 "github.com/tfkhdyt/openmusic-go/service/album"
)

// Injectors from wire.go:

func InitializeController() *Controller {
	config := postgres.NewConfig()
	db := postgres2.NewDB(config)
	repository := album.NewRepository(db)
	service := album2.NewService(repository)
	controller := NewController(service)
	return controller
}
