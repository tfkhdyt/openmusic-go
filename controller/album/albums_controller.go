package album

import (
	"fmt"
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

func (c Controller) Post(ctx *gin.Context) {
	var album albumEntity.Album
	fmt.Println("Lagi di controller")

	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	albumId, err := c.service.Create(album.Name, album.Year)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"albumId": albumId,
		},
	})
}

func (c Controller) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	album, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"album": album,
		},
	})
}

func (c Controller) Put(ctx *gin.Context) {
	id := ctx.Param("id")

	oldAlbum, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	var newAlbum albumEntity.Album

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err := c.service.Update(&oldAlbum, &newAlbum); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Album berhasil diubah",
	})
}

func (c Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	album, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err := c.service.Delete(&album); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Album berhasil dihapus",
	})
}
