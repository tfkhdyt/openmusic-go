package collab

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/collab"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Delete(ctx *gin.Context) {
	var collab collab.Collaboration

	// validate request body
	if err := ctx.ShouldBindJSON(&collab); err != nil {
		response.SendFailWithErrors(ctx, 400, err.Error())
		return
	}

	// get user id from middleware
	ownerId, ok := ctx.MustGet("userId").(string)
	if !ok {
		response.SendFail(ctx, 401, "User ID tidak valid")
		return
	}

	// verify playlist owner
	_, err := c.playlistsService.VerifyPlaylistOwner(collab.PlaylistId, ownerId)
	if err != nil {
		response.ErrorAssertion(ctx, err)
		return
	}

	// verify collaboration
	if err := c.collabsService.Verify(collab.PlaylistId, collab.UserId); err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	// delete collaboration
	if err := c.collabsService.Delete(&collab); err != nil {
		response.SendError(ctx, err)
		return
	}

	// success response
	response.SendSuccessWithMessage(ctx, 200, "Kolaborasi berhasil dihapus")
}
