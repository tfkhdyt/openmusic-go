package collab

import "github.com/tfkhdyt/openmusic-go/entity/collab"

func (r Repository) FindOne(playlistId string, userId string) (collab.Collaboration, error) {
	var collab collab.Collaboration

	err := r.db.First(&collab, "playlist_id = ? AND user_id = ?", playlistId, userId).Error

	return collab, err
}
