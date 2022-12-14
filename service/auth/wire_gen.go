// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package auth

import (
	"github.com/tfkhdyt/openmusic-go/repository/postgres/auth"
)

// Injectors from wire.go:

func InitializeService() *Service {
	repository := auth.NewRepository()
	service := NewService(repository)
	return service
}
