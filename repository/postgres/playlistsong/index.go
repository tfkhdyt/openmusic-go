package playlistsong

import (
	"github.com/tfkhdyt/openmusic-go/db/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	db := postgres.GetInstance()
	return &Repository{db}
}
