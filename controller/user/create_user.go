package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/user"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Create(ctx *gin.Context) {
	var user user.User

	// validate request body
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.SendFailWithErrors(ctx, 400, err.Error())
		return
	}

	// verify username to prevent already exist on database
	if err := c.service.VerifyNewUsername(user.Username); err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	// create user / register
	userId, err := c.service.Create(&user)
	if err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithData(ctx, 201, &gin.H{
		"userId": userId,
	})
}
