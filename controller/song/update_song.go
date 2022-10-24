package song

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/song"
)

func (c Controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	oldSong, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	var newSong song.Song

	if err := ctx.ShouldBindJSON(&newSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err := c.service.Update(&oldSong, &newSong); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Lagu berhasil diubah",
	})
}
