package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/db/postgres"
	"github.com/tfkhdyt/openmusic-go/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := postgres.InitializeDB()
	db.Connect()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	router := router.InitializeRouter()
	router.Route(r)

	if err := r.Run(fmt.Sprintf(":%v", port)); err != nil {
		panic(err.Error())
	}
}
