package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c Controller) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")

	album, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"album": album,
		},
	})
}