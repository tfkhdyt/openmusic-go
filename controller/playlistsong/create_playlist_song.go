package playlistsong

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/entity/playlistsong"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func (c Controller) Create(ctx *gin.Context) {
	var playlistSong playlistsong.PlaylistSong

	// validate request body
	if err := ctx.ShouldBindJSON(&playlistSong); err != nil {
		response.SendFailWithErrors(ctx, 400, err.Error())
		return
	}

	// get user id from middleware
	userId, ok := ctx.MustGet("userId").(string)
	if !ok {
		response.SendFail(ctx, 401, "User ID tidak valid")
		return
	}

	// get playlist id from param
	playlistId := ctx.Param("id")

	// verify playlist owner
	playlist, err := c.playlistsService.VerifyPlaylistAccess(playlistId, userId)
	if err != nil {
		response.ErrorAssertion(ctx, err)
		return
	}

	// find song
	song, err2 := c.songsService.FindOne(playlistSong.SongId)
	if err2 != nil {
		response.SendFail(ctx, err2.StatusCode, err2.Error())
		return
	}

	// add song to playlist
	if err := c.playlistSongsService.Create(&playlist, &song); err != nil {
		response.SendError(ctx, err)
		return
	}

	// add activity log
	if err := c.activitiesService.Create(playlistId, song.ID, userId, "add"); err != nil {
		response.SendError(ctx, err)
		return
	}

	// success response
	response.SendSuccessWithMessage(ctx, 201, "Lagu berhasil dimasukkan ke dalam playlist")
}
