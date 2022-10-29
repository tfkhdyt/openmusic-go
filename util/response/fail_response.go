package response

import (
	"github.com/gin-gonic/gin"
)

func SendFail(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"status":  "fail",
		"message": message,
	})
}
