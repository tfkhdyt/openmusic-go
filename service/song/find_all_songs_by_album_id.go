package song

import (
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) FindAllByAlbumID(id string) ([]song.Song, *exception.InternalServerError) {
	songs, err := s.repository.FindAllByAlbumID(id)
	if err != nil {
		return songs, exception.NewInternalServerError("Gagal mendapatkan lagu")
	}

	return songs, nil
}
