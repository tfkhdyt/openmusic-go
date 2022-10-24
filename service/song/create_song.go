package song

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Create(payload *song.Song) (string, *exception.InternalServerError) {
	id := fmt.Sprintf("song-%v", uuid.NewString())
	payload.ID = id

	songId, err := s.repository.Create(payload)
	if err != nil {
		return "", exception.NewInternalServerError("Gagal menambahkan lagu")
	}

	return songId, nil
}
