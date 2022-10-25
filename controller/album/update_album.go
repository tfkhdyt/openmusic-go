package album

import (
	"github.com/gin-gonic/gin"
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	oldAlbum, err := c.albumsService.FindOne(id)
	if err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	var newAlbum albumEntity.Album

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		response.SendFail(ctx, 400, err.Error())
		return
	}

	if err := c.albumsService.Update(&oldAlbum, &newAlbum); err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithMessage(ctx, "Album berhasil diubah")
}
