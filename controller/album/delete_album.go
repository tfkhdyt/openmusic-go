package album

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	album, err := c.albumsService.FindOne(id)
	if err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	if err := c.albumsService.Delete(&album); err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithMessage(ctx, "Album berhasil dihapus")
}
