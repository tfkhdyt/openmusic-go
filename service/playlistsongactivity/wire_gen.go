// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package playlistsongactivity

import (
	"github.com/tfkhdyt/openmusic-go/repository/postgres/playlistsongactivity"
)

// Injectors from wire.go:

func InitializeService() *Service {
	repository := playlistsongactivity.NewRepository()
	service := NewService(repository)
	return service
}
