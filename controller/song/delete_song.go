package song

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	song, err := c.service.FindOne(id)
	if err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	if err := c.service.Delete(&song); err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithMessage(ctx, 201, "Lagu berhasil dihapus")
}
