package playlist

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
)

func (r Repository) FindAll(userId string) ([]playlist.FindAllPlaylistsResult, error) {
	/*
		 	var user user.User

			err := r.db.Preload("Playlists").First(&user, "id = ?", userId).Error

			return user.Playlists, err
	*/
	result := make([]playlist.FindAllPlaylistsResult, 0)

	err := r.db.Table("playlists p").Select("p.id", "p.name", "u.username").Joins("JOIN users u ON u.id = p.user_id").Joins("JOIN user_playlists up ON up.playlist_id = p.id").Where("p.user_id = ? OR up.user_id = ?", userId, userId).Scan(&result).Error

	return result, err
}
