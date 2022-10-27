package playlist

import "github.com/tfkhdyt/openmusic-go/service/playlist"

type Controller struct {
	service *playlist.Service
}

func NewController(service *playlist.Service) *Controller {
	return &Controller{service}
}
