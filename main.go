package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/db/postgres"
	"github.com/tfkhdyt/openmusic-go/router/album"
	"github.com/tfkhdyt/openmusic-go/router/auth"
	"github.com/tfkhdyt/openmusic-go/router/song"
	"github.com/tfkhdyt/openmusic-go/router/user"
)

func main() {
	db := postgres.InitializeDB()
	db.Connect()
	r := gin.Default()

	albumRG := r.Group("/albums")
	albumRouter := album.InitializeRouter()
	albumRouter.Route(albumRG)

	songRG := r.Group("/songs")
	songRouter := song.InitializeRouter()
	songRouter.Route(songRG)

	userRG := r.Group("/users")
	userRouter := user.InitializeRouter()
	userRouter.Route(userRG)

	authRG := r.Group("/auth")
	authRouter := auth.InitializeRouter()
	authRouter.Route(authRG)

	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}
