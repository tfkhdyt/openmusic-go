package collab

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/collab"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Create(ctx *gin.Context) {
	var collab collab.Collaboration

	// validate request body
	if err := ctx.ShouldBindJSON(&collab); err != nil {
		response.SendFailWithErrors(ctx, 400, err.Error())
		return
	}

	ownerId, ok := ctx.MustGet("userId").(string)
	if !ok {
		response.SendFail(ctx, 401, "Owner ID tidak valid")
	}

	// verify playlist owner
	playlist, err := c.playlistsService.VerifyPlaylistOwner(collab.PlaylistId, ownerId)
	if err != nil {
		response.ErrorAssertion(ctx, err)
		return
	}

	// find user by id
	user, err2 := c.usersService.FindOne(collab.UserId)
	if err2 != nil {
		response.SendFail(ctx, err2.StatusCode, err2.Error())
		return
	}

	// add collab
	if err := c.collabsService.Create(&playlist, &user); err != nil {
		response.SendError(ctx, err)
		return
	}

	// success response
	response.SendSuccessWithMessage(ctx, 201, "Berhasil menambahkan collab")
}
