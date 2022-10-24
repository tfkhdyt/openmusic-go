package song

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c Controller) FindAll(ctx *gin.Context) {
	songs, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"songs": songs,
		},
	})
}
