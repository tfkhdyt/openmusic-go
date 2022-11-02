package playlist

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Delete(ctx *gin.Context) {
	// get user id from middleware
	userId, ok := ctx.MustGet("userId").(string)
	if !ok {
		response.SendFail(ctx, 401, "User ID tidak valid")
		return
	}

	// get playlist id from param
	playlistId := ctx.Param("id")

	// verify playlist owner
	playlist, err := c.service.VerifyPlaylistOwner(playlistId, userId)
	if err != nil {
		response.ErrorAssertion(ctx, err)
		return
	}

	// delete playlist
	if err := c.service.Delete(&playlist); err != nil {
		response.SendError(ctx, err)
		return
	}

	// success response
	response.SendSuccessWithMessage(ctx, 200, "Playlist berhasil dihapus")
}
