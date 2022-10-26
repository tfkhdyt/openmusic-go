package auth

import (
	"github.com/tfkhdyt/openmusic-go/entity/auth"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Store(auth *auth.Auth) *exception.InternalServerError {
	if err := s.repository.Store(auth); err != nil {
		return exception.NewInternalServerError("Gagal membuat refresh token")
	}

	return nil
}
