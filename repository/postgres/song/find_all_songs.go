package song

import (
	"github.com/tfkhdyt/openmusic-go/entity/song"
)

func (r Repository) FindAll(title string, performer string) ([]song.Song, error) {
	var songs []song.Song

	err := r.db.Where("title ILIKE ? OR performer ILIKE ?", title, performer).Select("id", "title", "performer").Find(&songs).Error

	return songs, err
}
