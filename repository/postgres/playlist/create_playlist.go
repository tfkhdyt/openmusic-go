package playlist

import "github.com/tfkhdyt/openmusic-go/entity/playlist"

func (r Repository) Create(playlist *playlist.Playlist) (string, error) {
	err := r.db.Create(playlist).Error

	return playlist.ID, err
}
