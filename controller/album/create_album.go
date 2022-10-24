package album

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
)

func (c Controller) Create(ctx *gin.Context) {
	var album albumEntity.Album

	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	albumId, err := c.albumsService.Create(album.Name, album.Year)
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
