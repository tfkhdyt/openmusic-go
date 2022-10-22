package album

import (
	"log"
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

func (a Controller) Post(c *gin.Context) {
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
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data": gin.H{
			"albumId": albumId,
		},
	})
}

func (a Controller) GetById(c *gin.Context) {
	id := c.Param("id")

	album, err := a.service.FindOne(id)
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

func (a Controller) Put(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum albumEntity.Album

	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err := a.service.Update(id, &updatedAlbum); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "album berhasil diubah",
	})
}
