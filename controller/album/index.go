package album

import (
	"github.com/tfkhdyt/openmusic-go/service/album"
	"github.com/tfkhdyt/openmusic-go/service/song"
)

type Controller struct {
	albumsService *album.Service
	songsService  *song.Service
}

func NewController(albumsService *album.Service, songsService *song.Service) *Controller {
	return &Controller{albumsService, songsService}
}
