package song

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/song"
)

func (c Controller) Create(ctx *gin.Context) {
	var song song.Song

	if err := ctx.ShouldBindJSON(&song); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	songId, err := c.service.Create(&song)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status": "error",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"songId": songId,
		},
	})
}
