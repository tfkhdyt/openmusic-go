package song

import "github.com/tfkhdyt/openmusic-go/entity/song"

func (r Repository) Create(payload *song.Song) (string, error) {
	result := r.db.Create(payload).Error

	return payload.ID, result
}
