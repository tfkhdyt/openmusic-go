package album

import "github.com/tfkhdyt/openmusic-go/entity/album"

func (r *Repository) Create(id string, name string, year uint16) (album.Album, error) {
	album := album.Album{
		ID:   id,
		Name: name,
		Year: year,
	}

	return album, r.db.Create(&album).Error
}
