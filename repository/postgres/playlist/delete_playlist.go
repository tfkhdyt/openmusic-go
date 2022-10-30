package playlist

import "github.com/tfkhdyt/openmusic-go/entity/playlist"

func (r Repository) Delete(playlist *playlist.Playlist) error {
	if err := r.db.Delete(playlist).Error; err != nil {
		return err
	}

	return nil
}
