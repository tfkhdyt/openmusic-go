package token

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tfkhdyt/openmusic-go/exception"
)

type CustomClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func (m Manager) GenerateJWT(userId string, expiredTime time.Duration) (string, *exception.InternalServerError) {
	claims := CustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiredTime).Unix(),
			Issuer:    "openmusic-go",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(m.config.GetSecretKey()))
	if err != nil {
		return tokenString, exception.NewInternalServerError("Gagal membuat token")
	}

	return tokenString, nil
}
