package user

import "github.com/tfkhdyt/openmusic-go/service/user"

type Controller struct {
	service *user.Service
}

func NewController(service *user.Service) *Controller {
	return &Controller{service}
}
