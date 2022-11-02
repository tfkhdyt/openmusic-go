package playlistsongactivity

import (
	"github.com/tfkhdyt/openmusic-go/service/playlist"
	"github.com/tfkhdyt/openmusic-go/service/playlistsongactivity"
)

type Controller struct {
	service          *playlistsongactivity.Service
	playlistsService *playlist.Service
}

func NewController(service *playlistsongactivity.Service, playlistsService *playlist.Service) *Controller {
	return &Controller{service, playlistsService}
}
