package album

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/album"
)

func SetAlbumsRoutes(router *gin.RouterGroup) {
	controller := album.InitializeController()

	router.POST("/", controller.PostHandler)
	router.GET("/:id", controller.GetByIdHandler)
}
