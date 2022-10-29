package jwt

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/tokenize/token"
	"github.com/tfkhdyt/openmusic-go/util/response"
)

func VerifyJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			response.SendFail(ctx, 401, "Token tidak boleh kosong")
		}

		tokenString := strings.Fields(header)

		if tokenString[0] != "Bearer" {
			response.SendFail(ctx, 401, "Jenis token harus berupa \"Bearer\"")
		}

		tokenManager := token.InitializeManager()

		userId, err := tokenManager.VerifyJWT(tokenString[1])
		if err != nil {
			response.SendFail(ctx, err.StatusCode, err.Error())
		}

		ctx.Set("userId", userId)

		ctx.Next()
	}
}
