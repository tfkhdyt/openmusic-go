package user

import (
	"github.com/tfkhdyt/openmusic-go/entity/user"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) FindByUsername(username string) (user.User, *exception.NotFoundError) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, exception.NewNotFoundError("User tidak ditemukan")
	}

	return user, nil
}
