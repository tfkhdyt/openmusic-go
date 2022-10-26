package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (m Manager) GenerateJWT(userId string, expiredTime time.Duration) (string, *exception.InternalServerError) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(expiredTime).Unix()
	claims["authorized"] = true
	claims["userId"] = userId

	secret := m.config.GetSecretKey()
	fmt.Println(secret)

	tokenString, err := token.SignedString([]byte(m.config.GetSecretKey()))
	if err != nil {
		return "", exception.NewInternalServerError("Gagal membuat token")
	}

	return tokenString, nil
}
