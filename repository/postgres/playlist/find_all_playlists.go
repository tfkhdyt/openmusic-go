package playlist

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/entity/user"
)

func (r Repository) FindAll(userId string) ([]playlist.Playlist, error) {
	var user user.User

	err := r.db.Preload("Playlists").First(&user, "id = ?", userId).Error

	return user.Playlists, err
}
