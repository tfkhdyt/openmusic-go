package song

import "github.com/tfkhdyt/openmusic-go/entity/song"

func (r Repository) Update(oldSong *song.Song, newSong *song.Song) error {
	err := r.db.Model(oldSong).Updates(newSong).Error

	return err
}
