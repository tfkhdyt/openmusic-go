package song

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")

	song, err := c.service.FindOne(id)
	if err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	response.SendSuccessWithData(ctx, 200, &gin.H{
		"song": song,
	})
}
