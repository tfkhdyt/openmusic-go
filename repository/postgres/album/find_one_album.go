package album

import "github.com/tfkhdyt/openmusic-go/entity/album"

func (r Repository) FindOne(id string) (album.Album, error) {
	var album album.Album
	err := r.db.First(&album, "id = ?", id).Error

	return album, err
}
