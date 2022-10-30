package playlistsong

import (
	"github.com/tfkhdyt/openmusic-go/service/playlist"
	"github.com/tfkhdyt/openmusic-go/service/playlistsong"
	"github.com/tfkhdyt/openmusic-go/service/song"
)

type Controller struct {
	playlistSongsService *playlistsong.Service
	playlistsService     *playlist.Service
	songsService         *song.Service
}

func NewController(playlistSongsService *playlistsong.Service, playlistsService *playlist.Service, songsService *song.Service) *Controller {
	return &Controller{playlistSongsService, playlistsService, songsService}
}
