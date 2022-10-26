package token

import (
	"github.com/golang-jwt/jwt"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (m Manager) VerifyJWT(tokenString string) (string, *exception.BadRequestError) {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(m.config.GetSecretKey()), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return "", exception.NewBadRequestError("Token tidak valid")
}
