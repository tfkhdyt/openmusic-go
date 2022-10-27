package playlist

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Create(ctx *gin.Context) {
	var playlist playlist.Playlist

	// validate request body
	if err := ctx.ShouldBindJSON(&playlist); err != nil {
		response.SendFail(ctx, 400, err.Error())
		return
	}

	// get userId from middleware
	userId := ctx.MustGet("userId").(string)
	playlist.UserId = userId

	// insert playlist to db
	playlistId, err := c.service.Create(&playlist)
	if err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithData(ctx, 201, &gin.H{
		"playlistId": playlistId,
	})
}
