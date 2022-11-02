package playlistsongactivity

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlistsongactivity"
)

func (r Repository) Create(activity *playlistsongactivity.PlaylistSongActivity) error {
	if err := r.db.Create(activity).Error; err != nil {
		return err
	}

	return nil
}
