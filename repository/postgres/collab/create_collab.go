package collab

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/entity/user"
)

func (r Repository) Create(playlist *playlist.Playlist, user *user.User) error {
	if err := r.db.Model(user).Association("Playlists").Append(playlist); err != nil {
		return err
	}

	return nil
}
