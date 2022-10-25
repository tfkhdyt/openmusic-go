package response

import (
	"github.com/gin-gonic/gin"
)

func SendSuccessWithMessage(ctx *gin.Context, message string) {
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": message,
	})
}
