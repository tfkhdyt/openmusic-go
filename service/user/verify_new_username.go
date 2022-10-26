package user

import "github.com/tfkhdyt/openmusic-go/exception"

func (s Service) VerifyNewUsername(username string) *exception.BadRequestError {
	_, err := s.repository.FindByUsername(username)
	if err != nil {
		return nil
	}

	return exception.NewBadRequestError("Username sudah digunakan")
}
