package album

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
)

func (c Controller) Create(ctx *gin.Context) {
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
