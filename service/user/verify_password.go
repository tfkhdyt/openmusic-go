package user

import (
	"github.com/tfkhdyt/openmusic-go/exception"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) VerifyPassword(hashedPassword string, password string) *exception.AuthenticationError {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return exception.NewAuthenticationError("Password tidak valid")
	}

	return nil
}
