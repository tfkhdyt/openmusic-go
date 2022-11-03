package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/router/album"
	"github.com/tfkhdyt/openmusic-go/router/auth"
	"github.com/tfkhdyt/openmusic-go/router/playlist"
	"github.com/tfkhdyt/openmusic-go/router/song"
	"github.com/tfkhdyt/openmusic-go/router/user"
)

func Route(r *gin.Engine) {
	albumRG := r.Group("/albums")
	albumRouter := album.InitializeRouter()
	albumRouter.Route(albumRG)

	songRG := r.Group("/songs")
	songRouter := song.InitializeRouter()
	songRouter.Route(songRG)

	userRG := r.Group("/users")
	userRouter := user.InitializeRouter()
	userRouter.Route(userRG)

	authRG := r.Group("/authentications")
	authRouter := auth.InitializeRouter()
	authRouter.Route(authRG)

	playlistRG := r.Group("/playlists")
	playlistRouter := playlist.InitializeRouter()
	playlistRouter.Route(playlistRG)
}
