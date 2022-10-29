package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/auth"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Login(ctx *gin.Context) {
	var credential auth.Credentials

	// validate request body
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		response.SendFailWithErrors(ctx, 400, err.Error())
		return
	}

	// check if user is exist
	user, err := c.usersService.FindByUsername(credential.Username)
	if err != nil {
		response.SendFail(ctx, 401, err.Error())
		return
	}

	// check password
	if err := c.usersService.VerifyPassword(user.Password, credential.Password); err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	// generate access token
	accessToken, err2 := c.tokenManager.GenerateJWT(user.ID, 5*time.Minute)
	if err2 != nil {
		response.SendError(ctx, err2)
		return
	}

	// generate refresh token
	refreshToken, err3 := c.tokenManager.GenerateJWT(user.ID, 730*time.Hour)
	if err3 != nil {
		response.SendError(ctx, err3)
		return
	}

	// store refresh token to db
	if err := c.authService.Store(&auth.Auth{Token: refreshToken}); err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithData(ctx, 201, &gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
