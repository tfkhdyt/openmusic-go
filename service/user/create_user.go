package user

import (
	"github.com/google/uuid"
	"github.com/tfkhdyt/openmusic-go/entity/user"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Create(user *user.User) (string, *exception.BadRequestError) {
	user.ID = "user-" + uuid.NewString()

	userId, err := s.repository.Create(user)
	if err != nil {
		return userId, exception.NewBadRequestError("Gagal membuat user baru")
	}

	return userId, nil
}
