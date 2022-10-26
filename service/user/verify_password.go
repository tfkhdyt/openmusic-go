package user

import (
	"github.com/tfkhdyt/openmusic-go/exception"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) VerifyPassword(id string, password string) *exception.AuthenticationError {
	user, _ := s.repository.FindOne(id)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return exception.NewAuthenticationError("Password tidak valid")
	}

	return nil
}
