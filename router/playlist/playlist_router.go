package playlist

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/playlist"
	"github.com/tfkhdyt/openmusic-go/controller/playlistsong"
	"github.com/tfkhdyt/openmusic-go/controller/playlistsongactivity"
	"github.com/tfkhdyt/openmusic-go/middleware/jwt"
)

type Router struct {
	playlistsController     *playlist.Controller
	playlistSongsController *playlistsong.Controller
	activitiesController    *playlistsongactivity.Controller
}

func NewRouter(playlistsController *playlist.Controller, playlistSongsController *playlistsong.Controller, activitiesController *playlistsongactivity.Controller) *Router {
	return &Router{playlistsController, playlistSongsController, activitiesController}
}

func (r Router) Route(router *gin.RouterGroup) {
	router.Use(jwt.VerifyJWT())

	router.POST("/", r.playlistsController.Create)
	router.GET("/", r.playlistsController.FindAll)
	router.DELETE("/:id", r.playlistsController.Delete)

	router.POST("/:id/songs", r.playlistSongsController.Create)
	router.GET("/:id/songs", r.playlistSongsController.FindAll)
	router.DELETE("/:id/songs", r.playlistSongsController.Delete)

	router.GET("/:id/activities", r.activitiesController.FindAll)
}
