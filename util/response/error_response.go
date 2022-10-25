package response

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func SendError(ctx *gin.Context, err *exception.InternalServerError) {
	ctx.JSON(err.StatusCode, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
	log.Println(err)
}
