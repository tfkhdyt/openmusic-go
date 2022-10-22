package album

import "github.com/tfkhdyt/openmusic-go/entity/album"

func (r Repository) Delete(album *album.Album) error {
	err := r.db.Delete(album).Error

	return err
}
