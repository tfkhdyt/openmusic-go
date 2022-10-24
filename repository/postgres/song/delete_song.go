package song

import "github.com/tfkhdyt/openmusic-go/entity/song"

func (r Repository) Delete(song *song.Song) error {
	err := r.db.Delete(song).Error

	return err
}
