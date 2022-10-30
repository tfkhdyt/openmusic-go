package playlist

import "github.com/tfkhdyt/openmusic-go/entity/playlist"

func (r Repository) FindOne(id string) (playlist.Playlist, error) {
	var playlist playlist.Playlist

	err := r.db.First(&playlist, "id = ?", id).Error

	return playlist, err
}
