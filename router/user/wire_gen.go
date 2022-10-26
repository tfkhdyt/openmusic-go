// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user

import (
	user2 "github.com/tfkhdyt/openmusic-go/controller/user"
	"github.com/tfkhdyt/openmusic-go/service/user"
)

// Injectors from wire.go:

func InitializeRouter() *Router {
	service := user.InitializeService()
	controller := user2.NewController(service)
	router := NewRouter(controller)
	return router
}
