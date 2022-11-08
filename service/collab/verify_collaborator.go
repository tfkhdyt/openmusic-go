package collab

import (
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Verify(playlistId string, userId string) *exception.AuthorizationError {
	_, err := s.repository.FindOne(playlistId, userId)
	if err != nil {
		return exception.NewAuthorizationError("Kolaborasi gagal diverifikasi")
	}

	return nil
}
