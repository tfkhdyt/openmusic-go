package collab

import "github.com/tfkhdyt/openmusic-go/entity/collab"

func (r Repository) Delete(collab *collab.Collaboration) error {
	err := r.db.Delete(collab, "playlist_id = ? AND user_id = ?", collab.PlaylistId, collab.UserId).Error

	return err
}
