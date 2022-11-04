//go:build wireinject
// +build wireinject

package router

import (
	"github.com/google/wire"
	albumsController "github.com/tfkhdyt/openmusic-go/controller/album"
	authsController "github.com/tfkhdyt/openmusic-go/controller/auth"
	collabsController "github.com/tfkhdyt/openmusic-go/controller/collab"
	playlistsController "github.com/tfkhdyt/openmusic-go/controller/playlist"
	playlistSongsController "github.com/tfkhdyt/openmusic-go/controller/playlistsong"
	playlistSongActivity "github.com/tfkhdyt/openmusic-go/controller/playlistsongactivity"
	songsController "github.com/tfkhdyt/openmusic-go/controller/song"
	usersController "github.com/tfkhdyt/openmusic-go/controller/user"
	albumsService "github.com/tfkhdyt/openmusic-go/service/album"
	authsService "github.com/tfkhdyt/openmusic-go/service/auth"
	"github.com/tfkhdyt/openmusic-go/service/collab"
	playlistsService "github.com/tfkhdyt/openmusic-go/service/playlist"
	"github.com/tfkhdyt/openmusic-go/service/playlistsong"
	playlistSongActivities "github.com/tfkhdyt/openmusic-go/service/playlistsongactivity"
	songsService "github.com/tfkhdyt/openmusic-go/service/song"
	usersService "github.com/tfkhdyt/openmusic-go/service/user"
	"github.com/tfkhdyt/openmusic-go/tokenize/token"
)

func InitializeRouter() *Router {
	wire.Build(
		NewRouter,

		// albums
		albumsController.NewController,
		albumsService.InitializeService,
		songsService.InitializeService,

		// songs
		songsController.NewController,

		// users
		usersController.NewController,
		usersService.InitializeService,

		// auths
		authsController.NewController,
		authsService.InitializeService,
		token.InitializeManager,

		// playlists
		playlistsController.NewController,
		playlistsService.InitializeService,
		playlistSongActivities.InitializeService,

		// playlist songs
		playlistSongsController.NewController,
		playlistsong.InitializeService,

		// playlist song activities
		playlistSongActivity.NewController,

		// collabs
		collabsController.NewController,
		collab.InitializeService,
	)

	return nil
}
