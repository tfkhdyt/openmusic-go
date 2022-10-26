package auth

import "github.com/tfkhdyt/openmusic-go/exception"

func (s Service) VerifyRefreshToken(token string) *exception.BadRequestError {
	if err := s.repository.FindOne(token); err != nil {
		return exception.NewBadRequestError("Refresh token tidak terdaftar di database")
	}

	return nil
}
