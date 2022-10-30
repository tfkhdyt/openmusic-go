package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/auth"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) DeleteToken(ctx *gin.Context) {
	var auth auth.Auth

	// validate request body
	if err := ctx.ShouldBindJSON(&auth); err != nil {
		response.SendFailWithErrors(ctx, 400, err.Error())
		return
	}

	// check token in database
	if err := c.authService.VerifyRefreshToken(auth.Token); err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	// delete token from database
	if err := c.authService.DeleteToken(auth.Token); err != nil {
		response.SendError(ctx, err)
		return
	}

	// send success response
	response.SendSuccessWithMessage(ctx, 200, "Token berhasil dihapus")
}
