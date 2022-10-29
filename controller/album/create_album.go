package album

import (
	"github.com/gin-gonic/gin"
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Create(ctx *gin.Context) {
	var album albumEntity.Album

	if err := ctx.ShouldBindJSON(&album); err != nil {
		response.SendFailWithErrors(ctx, 400, err.Error())
		return
	}

	albumId, err := c.albumsService.Create(album.Name, album.Year)
	if err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithData(ctx, 201, &gin.H{
		"albumId": albumId,
	})
}
