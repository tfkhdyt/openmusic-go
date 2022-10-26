package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/auth"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) RefreshToken(ctx *gin.Context) {
	var auth auth.Auth

	// validate request body
	if err := ctx.ShouldBindJSON(&auth); err != nil {
		response.SendFail(ctx, 400, err.Error())
		return
	}

	// check token on database
	if err := c.authService.VerifyRefreshToken(auth.Token); err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	// verify token
	userId, err := c.tokenManager.VerifyJWT(auth.Token)
	if err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	// generate access token
	accessToken, err2 := c.tokenManager.GenerateJWT(userId, 5*time.Minute)
	if err2 != nil {
		response.SendError(ctx, err2)
		return
	}

	response.SendSuccessWithData(ctx, 200, &gin.H{
		"accessToken": accessToken,
	})
}
