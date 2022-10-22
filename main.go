package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/router/album"
)

func main() {
	r := gin.Default()

	albumsRouterGroup := r.Group("/albums")
	album.SetAlbumsRoutes(albumsRouterGroup)

	r.Run()
}
