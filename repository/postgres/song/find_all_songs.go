package song

import "github.com/tfkhdyt/openmusic-go/entity/song"

func (r Repository) FindAll() ([]song.Song, error) {
	var songs []song.Song
	err := r.db.Select("id", "title", "performer").Find(&songs).Error

	return songs, err
}
