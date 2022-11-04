package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/db/postgres"
	"github.com/tfkhdyt/openmusic-go/router"
)

func main() {
	db := postgres.InitializeDB()
	db.Connect()

	r := gin.Default()
	router := router.InitializeRouter()
	router.Route(r)

	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}
