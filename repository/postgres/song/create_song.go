package song

import "github.com/tfkhdyt/openmusic-go/entity/song"

func (r Repository) Create(song *song.Song) (string, error) {
	err := r.db.Create(song).Error

	return song.ID, err
}
