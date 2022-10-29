package playlist

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) FindAll(ctx *gin.Context) {
	userId, ok := ctx.MustGet("userId").(string)
	if !ok {
		response.SendFail(ctx, 401, "User ID tidak valid")
		return
	}

	playlists, err := c.service.FindAll(userId)
	if err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	response.SendSuccessWithData(ctx, 200, &gin.H{
		"playlists": playlists,
	})
}
