package playlistsongactivity

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlistsongactivity"
)

func (r Repository) FindAll(id string) ([]playlistsongactivity.FindAllActivities, error) {
	result := make([]playlistsongactivity.FindAllActivities, 0)

	// err := r.db.Preload("Activities").First(&playlist, "id = ?", id).Error
	err := r.db.Table("playlist_song_activities psa").Select("u.username", "s.title", "psa.action", "psa.time").Joins("JOIN users u ON u.id = psa.user_id").Joins("JOIN songs s ON s.id = psa.song_id").Order("psa.time").Where("psa.playlist_id = ?", id).Scan(&result).Error

	return result, err
}
