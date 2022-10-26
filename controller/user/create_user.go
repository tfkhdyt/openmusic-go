package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/user"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Create(ctx *gin.Context) {
	var user user.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.SendFail(ctx, 400, err.Error())
		return
	}

	if err := c.service.VerifyNewUsername(user.Username); err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	userId, err := c.service.Create(&user)
	if err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	response.SendSuccessWithData(ctx, 201, &gin.H{
		"userId": userId,
	})
}
