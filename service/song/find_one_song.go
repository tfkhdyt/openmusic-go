package song

import (
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) FindOne(id string) (song.Song, *exception.NotFoundError) {
	song, err := s.repository.FindOne(id)
	if err != nil {
		return song, exception.NewNotFoundError("Lagu tidak ditemukan")
	}

	return song, nil
}
