package song

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	song, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err := c.service.Delete(&song); err != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Lagu berhasil dihapus",
	})
}
