package song

import "github.com/tfkhdyt/openmusic-go/entity/song"

func (r Repository) FindAllByAlbumID(id string) ([]song.Song, error) {
	var songs []song.Song
	err := r.db.Select("id", "title", "performer").Find(&songs, "album_id = ?", id).Error

	return songs, err
}
