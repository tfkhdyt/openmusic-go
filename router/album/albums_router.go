package album

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/album"
)

type Router struct {
	controller *album.Controller
}

func NewRouter(controller *album.Controller) *Router {
	return &Router{controller}
}

func (r Router) Route(router *gin.RouterGroup) {

	router.POST("/", r.controller.Post)
	router.GET("/:id", r.controller.GetById)
	router.PUT("/:id", r.controller.Put)
	router.DELETE("/:id", r.controller.Delete)
}
