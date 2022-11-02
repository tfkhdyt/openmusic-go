package playlistsong

import (
	"github.com/tfkhdyt/openmusic-go/service/playlist"
	"github.com/tfkhdyt/openmusic-go/service/playlistsong"
	"github.com/tfkhdyt/openmusic-go/service/playlistsongactivity"
	"github.com/tfkhdyt/openmusic-go/service/song"
)

type Controller struct {
	playlistSongsService *playlistsong.Service
	playlistsService     *playlist.Service
	songsService         *song.Service
	activitiesService    *playlistsongactivity.Service
}

func NewController(playlistSongsService *playlistsong.Service, playlistsService *playlist.Service, songsService *song.Service, activitiesService *playlistsongactivity.Service) *Controller {
	return &Controller{playlistSongsService, playlistsService, songsService, activitiesService}
}
