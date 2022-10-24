package song

import "github.com/tfkhdyt/openmusic-go/service/song"

type Controller struct {
	service *song.Service
}

func NewController(service *song.Service) *Controller {
	return &Controller{service}
}
