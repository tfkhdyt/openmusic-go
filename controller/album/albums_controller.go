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
	var album albumEntity.Album

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	albumId, err := a.service.Create(album.Name, album.Year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data": gin.H{
			"albumId": albumId,
		},
	})
}

func (a Controller) GetByIdHandler(c *gin.Context) {
	albumId := c.Param("id")

	album, err := a.service.FindOne(albumId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": gin.H{
			"album": album,
		},
	})
}
