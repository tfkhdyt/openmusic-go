package album

import (
	albumService "github.com/tfkhdyt/openmusic-go/service/album"
)

type Controller struct {
	service *albumService.Service
}

func NewController(service *albumService.Service) *Controller {
	return &Controller{service}
}
