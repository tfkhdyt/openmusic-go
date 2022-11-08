package playlist

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) VerifyPlaylistAccess(playlistId string, userId string) (playlist.Playlist, any) {
	playlist, err := s.VerifyPlaylistOwner(playlistId, userId)
	if err != nil {
		switch e := err.(type) {
		case *exception.NotFoundError:
			return playlist, e
		default:
			err2 := s.collabsService.Verify(playlistId, userId)
			if err2 != nil {
				return playlist, err2
			}
		}
	}

	return playlist, nil
}
