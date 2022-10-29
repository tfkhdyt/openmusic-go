package song

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Create(ctx *gin.Context) {
	var song song.Song

	if err := ctx.ShouldBindJSON(&song); err != nil {
		response.SendFailWithErrors(ctx, 400, err.Error())
		return
	}

	songId, err := c.service.Create(&song)
	if err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithData(ctx, 201, &gin.H{
		"songId": songId,
	})
}
