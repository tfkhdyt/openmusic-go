package postgres

import (
	"fmt"

	postgresConfig "github.com/tfkhdyt/openmusic-go/config/postgres"
	"github.com/tfkhdyt/openmusic-go/entity/album"
	"github.com/tfkhdyt/openmusic-go/entity/auth"
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/entity/user"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgresDB *gorm.DB
)

type DB struct {
	config *postgresConfig.Config
}

func NewDB(config *postgresConfig.Config) *DB {
	return &DB{config: config}
}

func (d *DB) Connect() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", d.config.GetHost(), d.config.GetUser(), d.config.GetPass(), d.config.GetDBName(), d.config.GetPort())

	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&album.Album{}, &song.Song{}, &user.User{}, &auth.Auth{})
	PostgresDB = db

	fmt.Println("Connected to DB...")
}
