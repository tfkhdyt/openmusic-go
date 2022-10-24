package song

import "github.com/tfkhdyt/openmusic-go/entity/song"

func (r Repository) FindOne(id string) (song.Song, error) {
	var song song.Song

	err := r.db.First(&song, "id = ?", id).Error

	return song, err
}
