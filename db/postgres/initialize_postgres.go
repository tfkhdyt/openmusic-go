package postgres

import (
	"fmt"

	"github.com/tfkhdyt/openmusic-go/config/postgres"
	"github.com/tfkhdyt/openmusic-go/entity/album"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewDB(config *postgres.Config) *DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", config.Host, config.Username, config.Password, config.DbName, config.Port)

	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&album.Album{})

	return &DB{db}
}
