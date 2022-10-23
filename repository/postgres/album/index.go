package album

import (
	"github.com/tfkhdyt/openmusic-go/db/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	db := postgres.AlbumDB
	return &Repository{db}
}
