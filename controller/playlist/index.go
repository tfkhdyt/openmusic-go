package playlist

import (
	"github.com/tfkhdyt/openmusic-go/service/playlist"
	"github.com/tfkhdyt/openmusic-go/service/playlistsongactivity"
)

type Controller struct {
	playlistsService  *playlist.Service
	activitiesService *playlistsongactivity.Service
}

func NewController(playlistsService *playlist.Service, activitiesService *playlistsongactivity.Service) *Controller {
	return &Controller{playlistsService, activitiesService}
}
