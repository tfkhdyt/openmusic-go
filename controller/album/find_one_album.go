package album

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")

	album, err := c.albumsService.FindOne(id)
	if err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	songs, err2 := c.songsService.FindAllByAlbumID(album.ID)
	if err2 != nil {
		response.SendError(ctx, err2)
		return
	}

	album.Songs = songs

	response.SendSuccessWithData(ctx, 200, &gin.H{
		"album": album,
	})
}
