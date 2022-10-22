package postgres

import (
	"fmt"

	postgresConfig "github.com/tfkhdyt/openmusic-go/config/postgres"
	"github.com/tfkhdyt/openmusic-go/entity/album"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewDB() *DB {
	config := postgresConfig.Config{}
	dbHost := config.GetHost()
	dbUser := config.GetUser()
	dbPass := config.GetPassword()
	dbName := config.GetDBName()
	dbPort := config.GetPort()

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&album.Album{})

	return &DB{db}
}
