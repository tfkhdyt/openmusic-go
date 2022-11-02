package playlist

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) VerifyPlaylistOwner(playlistId string, userId string) (playlist.Playlist, any) {
	playlist, err := s.repository.FindOne(playlistId)
	if err != nil {
		return playlist, exception.NewNotFoundError("Playlist tidak ditemukan")
	}

	if playlist.UserId != userId {
		return playlist, exception.NewAuthorizationError("Anda tidak berhak mengakses playlist ini")
	}

	return playlist, nil
}
