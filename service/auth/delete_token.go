package auth

import (
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) DeleteToken(token string) *exception.InternalServerError {
	if err := s.repository.DeleteToken(token); err != nil {
		return exception.NewInternalServerError("Gagal menghapus token")
	}

	return nil
}
