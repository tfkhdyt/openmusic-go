package album

import (
	"github.com/tfkhdyt/openmusic-go/db/postgres"
	"github.com/tfkhdyt/openmusic-go/entity/album"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *postgres.DB) *Repository {
	return &Repository{db.DB}
}

func (r Repository) CreateAlbum(id string, name string, year uint16) (album.Album, error) {
	album := album.Album{
		ID:   id,
		Name: name,
		Year: year,
	}

	return album, r.DB.Create(&album).Error
}
