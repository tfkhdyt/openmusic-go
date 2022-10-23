package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/db/postgres"
	"github.com/tfkhdyt/openmusic-go/router/album"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	DB = postgres.InitializeDB().Connect()
}

func main() {
	r := gin.Default()

	albumRouter := album.InitializeRouter()
	albumRouter.Route(r.Group("/albums"))

	r.Run()
}
