package playlistsongactivity

import (
	"time"

	"github.com/google/uuid"
	"github.com/tfkhdyt/openmusic-go/entity/playlistsongactivity"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Create(playlistId string, songId string, userId string, action string) *exception.InternalServerError {
	id := "activity" + uuid.NewString()
	time := time.Now()

	activity := playlistsongactivity.PlaylistSongActivity{
		ID:         id,
		PlaylistID: playlistId,
		SongID:     songId,
		UserID:     userId,
		Action:     action,
		Time:       time,
	}

	if err := s.repository.Create(&activity); err != nil {
		return exception.NewInternalServerError("Gagal menambah activity")
	}

	return nil
}
