package album

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c Controller) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")

	album, err := c.albumsService.FindOne(id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	songs, err2 := c.songsService.FindAllByAlbumID(album.ID)
	if err2 != nil {
		ctx.JSON(err2.StatusCode, gin.H{
			"status":  "error",
			"message": err2.Error(),
		})
		log.Println(err2)
		return
	}

	album.Songs = songs

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"album": album,
		},
	})
}
