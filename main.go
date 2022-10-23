package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/db/postgres"
	"github.com/tfkhdyt/openmusic-go/router/album"
)

func main() {
	db := postgres.InitializeDB()
	db.Connect()
	r := gin.Default()

	albumRG := r.Group("/albums")
	albumRouter := album.InitializeRouter()
	albumRouter.Route(albumRG)

	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}
