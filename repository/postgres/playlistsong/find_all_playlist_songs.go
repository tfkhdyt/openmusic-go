package playlistsong

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"gorm.io/gorm"
)

func (r Repository) FindAll(id string) (playlist.Playlist, error) {
	var playlist playlist.Playlist
	err := r.db.Preload("Songs", func(tx *gorm.DB) *gorm.DB {
		return tx.Omit("year", "genre", "duration", "album_id")
	}).First(&playlist, "id = ?", id).Error

	// err := r.db.Model(&playlist.Playlist{}).Select("playlists.id, playlists.name, users.username").Joins("JOIN users on users.id = playlists.user_id").Joins("JOIN playlist_songs ON playlist_songs.playlist_id = ?", id).Where("playlists.id = ?", id).Scan(&playlist).Error

	return playlist, err
}
