package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
	albumService "github.com/tfkhdyt/openmusic-go/service/album"
)

type Controller struct {
	service *albumService.Service
}

func NewController(service *albumService.Service) *Controller {
	return &Controller{service}
}

func (a Controller) PostHandler(c *gin.Context) {
	var album albumEntity.Entity

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	a.service.AddAlbum(album.Name, album.Year)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
