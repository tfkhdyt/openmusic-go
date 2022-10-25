package response

import "github.com/gin-gonic/gin"

func SendSuccessWithData(ctx *gin.Context, statusCode int, data *gin.H) {
	ctx.JSON(statusCode, gin.H{
		"status": "success",
		"data":   data,
	})
}
