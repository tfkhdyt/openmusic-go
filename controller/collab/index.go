package collab

import (
	"github.com/tfkhdyt/openmusic-go/service/collab"
	"github.com/tfkhdyt/openmusic-go/service/playlist"
	"github.com/tfkhdyt/openmusic-go/service/user"
)

type Controller struct {
	collabsService   *collab.Service
	playlistsService *playlist.Service
	usersService     *user.Service
}

func NewController(collabsService *collab.Service, playlistsService *playlist.Service, usersService *user.Service) *Controller {
	return &Controller{collabsService, playlistsService, usersService}
}
