package song

import (
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Update(oldSong *song.Song, newSong *song.Song) *exception.InternalServerError {
	if err := s.repository.Update(oldSong, newSong); err != nil {
		return exception.NewInternalServerError("Gagal mengubah lagu")
	}

	return nil
}
