package playlist

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/playlist"
	"github.com/tfkhdyt/openmusic-go/middleware/jwt"
)

type Router struct {
	controller *playlist.Controller
}

func NewRouter(controller *playlist.Controller) *Router {
	return &Router{controller}
}

func (r Router) Route(router *gin.RouterGroup) {
	router.Use(jwt.VerifyJWT())
	router.POST("/", r.controller.Create)
}
