package album

import "github.com/tfkhdyt/openmusic-go/entity/album"

func (r Repository) Update(album *album.Album, updatedAlbum *album.Album) error {
	err := r.db.Model(album).Updates(updatedAlbum).Error

	return err
}
