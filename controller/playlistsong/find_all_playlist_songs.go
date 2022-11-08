package playlistsong

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) FindAll(ctx *gin.Context) {
	// get user id from middleware
	userId, ok := ctx.MustGet("userId").(string)
	if !ok {
		response.SendFail(ctx, 401, "User ID tidak valid")
		return
	}

	// get playlist id from param
	playlistId := ctx.Param("id")

	// verify playlist owner
	_, err := c.playlistsService.VerifyPlaylistAccess(playlistId, userId)
	if err != nil {
		response.ErrorAssertion(ctx, err)
		return
	}

	// find all playlist songs
	playlist, err2 := c.playlistSongsService.FindAll(playlistId)
	if err != nil {
		response.SendFail(ctx, err2.StatusCode, err2.Error())
		return
	}

	// success response
	response.SendSuccessWithData(ctx, 200, &gin.H{
		"playlist": playlist,
	})
}
