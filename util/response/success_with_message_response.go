package response

import (
	"github.com/gin-gonic/gin"
)

func SendSuccessWithMessage(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"status":  "success",
		"message": message,
	})
}
