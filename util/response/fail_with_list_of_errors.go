package response

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func SendFailWithErrors(ctx *gin.Context, statusCode int, message string) {
	errors := strings.Split(message, "\n")

	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"status": "fail",
		"errors": errors,
	})
}
