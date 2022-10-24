package song

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/song"
)

type Router struct {
	controller *song.Controller
}

func NewRouter(controller *song.Controller) *Router {
	return &Router{controller}
}

func (r Router) Route(router *gin.RouterGroup) {
	router.POST("/", r.controller.Create)
}
