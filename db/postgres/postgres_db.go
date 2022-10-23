package postgres

import (
	"fmt"

	postgresConfig "github.com/tfkhdyt/openmusic-go/config/postgres"
	"github.com/tfkhdyt/openmusic-go/entity/album"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	config *postgresConfig.Config
}

func NewDB(config *postgresConfig.Config) *DB {
	return &DB{config}
}

func (d DB) Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", d.config.Host, d.config.Username, d.config.Password, d.config.DbName, d.config.Port)

	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&album.Album{})

	return db
}
