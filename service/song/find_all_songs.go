package song

import (
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) FindAll() ([]song.Song, *exception.InternalServerError) {
	songs, err := s.repository.FindAll()
	if err != nil {
		return songs, exception.NewInternalServerError("Gagal mendapatkan lagu")
	}

	return songs, nil
}
