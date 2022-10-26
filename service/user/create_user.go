package user

import (
	"github.com/google/uuid"
	"github.com/tfkhdyt/openmusic-go/entity/user"
	"github.com/tfkhdyt/openmusic-go/exception"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Create(user *user.User) (string, *exception.InternalServerError) {
	user.ID = "user-" + uuid.NewString()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "", exception.NewInternalServerError("Gagal mengenkripsi password")
	}
	user.Password = string(hashedPassword)

	userId, err := s.repository.Create(user)
	if err != nil {
		return userId, exception.NewInternalServerError("Gagal membuat user baru")
	}

	return userId, nil
}
