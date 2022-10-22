package album

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/album"
)

func SetAlbumsRoutes(router *gin.RouterGroup) {
	controller := album.InitializeController()

	router.POST("/", controller.Post)
	router.GET("/:id", controller.GetById)
	router.PUT("/:id", controller.Put)
	router.DELETE("/:id", controller.Delete)
}
