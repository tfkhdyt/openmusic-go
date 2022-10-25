package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccessWithMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}
