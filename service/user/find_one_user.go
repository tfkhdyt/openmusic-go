package user

import (
	"github.com/tfkhdyt/openmusic-go/entity/user"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) FindOne(userId string) (user.User, *exception.NotFoundError) {
	user, err := s.repository.FindOne(userId)
	if err != nil {
		return user, exception.NewNotFoundError("User tidak ditemukan")
	}

	return user, nil
}
