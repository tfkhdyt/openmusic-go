package song

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	oldSong, err := c.service.FindOne(id)
	if err != nil {
		response.SendFail(ctx, err.StatusCode, err.Error())
		return
	}

	var newSong song.Song

	if err := ctx.ShouldBindJSON(&newSong); err != nil {
		response.SendFailWithErrors(ctx, 400, err.Error())
		return
	}

	if err := c.service.Update(&oldSong, &newSong); err != nil {
		response.SendError(ctx, err)
		return
	}

	response.SendSuccessWithMessage(ctx, 201, "Lagu berhasil diubah")
}
