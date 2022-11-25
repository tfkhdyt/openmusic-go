package postgres

import (
	"fmt"
	"log"

	postgresConfig "github.com/tfkhdyt/openmusic-go/config/postgres"
	"github.com/tfkhdyt/openmusic-go/entity/album"
	"github.com/tfkhdyt/openmusic-go/entity/auth"
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/entity/playlistsongactivity"
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/entity/user"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgresDB *gorm.DB
)

func init() {
	config := postgresConfig.NewConfig()

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", config.GetHost(), config.GetUser(), config.GetPass(), config.GetDBName(), config.GetPort())

	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&album.Album{}, &song.Song{}, &user.User{}, &auth.Auth{}, &playlist.Playlist{}, &playlistsongactivity.PlaylistSongActivity{})
	postgresDB = db

	log.Println("Connected to DB")
}

func GetInstance() *gorm.DB {
	return postgresDB
}

