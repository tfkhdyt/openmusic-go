package playlist

import (
	"github.com/google/uuid"
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Create(playlist *playlist.Playlist) (string, *exception.InternalServerError) {
	playlist.ID = "playlist-" + uuid.NewString()

	playlistId, err := s.repository.Create(playlist)
	if err != nil {
		return playlistId, exception.NewInternalServerError("Gagal membuat playlist")
	}

	return playlistId, nil
}
