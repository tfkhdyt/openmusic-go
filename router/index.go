package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/album"
	"github.com/tfkhdyt/openmusic-go/controller/auth"
	"github.com/tfkhdyt/openmusic-go/controller/collab"
	"github.com/tfkhdyt/openmusic-go/controller/playlist"
	"github.com/tfkhdyt/openmusic-go/controller/playlistsong"
	"github.com/tfkhdyt/openmusic-go/controller/playlistsongactivity"
	"github.com/tfkhdyt/openmusic-go/controller/song"
	"github.com/tfkhdyt/openmusic-go/controller/user"
	"github.com/tfkhdyt/openmusic-go/middleware/jwt"
)

type Router struct {
	albumsController        *album.Controller
	songsController         *song.Controller
	usersController         *user.Controller
	authsController         *auth.Controller
	playlistsController     *playlist.Controller
	playlistSongsController *playlistsong.Controller
	activitiesController    *playlistsongactivity.Controller
	collabsController       *collab.Controller
}

func NewRouter(albumsController *album.Controller, songsController *song.Controller, usersController *user.Controller, authsController *auth.Controller, playlistsController *playlist.Controller, playlistSongsController *playlistsong.Controller, activitiesController *playlistsongactivity.Controller, collabsController *collab.Controller) *Router {
	return &Router{albumsController, songsController, usersController, authsController, playlistsController, playlistSongsController, activitiesController, collabsController}
}

func (r Router) Route(router *gin.Engine) {
	// albums
	router.POST("/albums/", r.albumsController.Create)
	router.GET("/albums/:id", r.albumsController.FindOne)
	router.PUT("/albums/:id", r.albumsController.Update)
	router.DELETE("/albums/:id", r.albumsController.Delete)

	// songs
	router.POST("/songs/", r.songsController.Create)
	router.GET("/songs/", r.songsController.FindAll)
	router.GET("/songs/:id", r.songsController.FindOne)
	router.PUT("/songs/:id", r.songsController.Update)
	router.DELETE("/songs/:id", r.songsController.Delete)

	// users
	router.POST("/users", r.usersController.Create)

	// auths
	router.POST("/authentications", r.authsController.Login)
	router.PUT("/authentications", r.authsController.RefreshToken)
	router.DELETE("/authentications", r.authsController.DeleteToken)

	// playlists
	router.POST("/playlists/", jwt.VerifyJWT(), r.playlistsController.Create)
	router.GET("/playlists/", jwt.VerifyJWT(), r.playlistsController.FindAll)
	router.DELETE("/playlists/:id", jwt.VerifyJWT(), r.playlistsController.Delete)

	// playlist's songs
	router.POST("/playlists/:id/songs", jwt.VerifyJWT(), r.playlistSongsController.Create)
	router.GET("/playlists/:id/songs", jwt.VerifyJWT(), r.playlistSongsController.FindAll)
	router.DELETE("/playlists/:id/songs", jwt.VerifyJWT(), r.playlistSongsController.Delete)

	// playlist's activities
	router.GET("/playlists/:id/activities", jwt.VerifyJWT(), r.activitiesController.FindAll)

	// collabs
	router.POST("/collaborations", jwt.VerifyJWT(), r.collabsController.Create)
	router.DELETE("/collaborations", jwt.VerifyJWT(), r.collabsController.Delete)
}
