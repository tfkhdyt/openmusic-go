package album

import (
	"github.com/tfkhdyt/openmusic-go/db/postgres"
	"github.com/tfkhdyt/openmusic-go/entity/album"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *postgres.DB) *Repository {
	return &Repository{db.DB}
}

func (r Repository) Create(id string, name string, year uint16) (album.Album, error) {
	album := album.Album{
		ID:   id,
		Name: name,
		Year: year,
	}

	return album, r.db.Create(&album).Error
}

func (r Repository) FindOne(id string) (album.Album, error) {
	var album album.Album
	err := r.db.First(&album, "id = ?", id).Error

	return album, err
}

func (r Repository) Update(album *album.Album, updatedAlbum *album.Album) error {
	err := r.db.Model(album).Updates(updatedAlbum).Error

	return err
}

func (r Repository) Delete(album *album.Album) error {
	err := r.db.Delete(album).Error

	return err
}
