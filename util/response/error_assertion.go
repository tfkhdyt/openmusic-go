package response

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func ErrorAssertion(ctx *gin.Context, err any) {
	switch e := err.(type) {
	case *exception.NotFoundError:
		SendFail(ctx, e.StatusCode, e.Error())
	case *exception.AuthorizationError:
		SendFail(ctx, e.StatusCode, e.Error())
	}
}
