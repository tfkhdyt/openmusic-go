package playlistsong

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) FindAll(id string) (playlist.Playlist, *exception.NotFoundError) {
	playlist, err := s.repository.FindAll(id)
	if err != nil {
		return playlist, exception.NewNotFoundError("Playlist tidak ditemukan")
	}

	return playlist, nil
}
