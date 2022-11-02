package playlist

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) FindAll(userId string) ([]playlist.FindAllPlaylistsResult, *exception.NotFoundError) {
	playlists, err := s.repository.FindAll(userId)
	if err != nil {
		return playlists, exception.NewNotFoundError("Playlist tidak ditemukan")
	}

	return playlists, nil
}
