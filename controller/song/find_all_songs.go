package song

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) FindAll(ctx *gin.Context) {
	title := ctx.Query("title")
	performer := ctx.Query("performer")

	songs, err := c.service.FindAll(title, performer)
	if err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithData(ctx, 200, &gin.H{
		"songs": songs,
	})
}
