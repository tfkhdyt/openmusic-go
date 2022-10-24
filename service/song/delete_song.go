package song

import (
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Delete(song *song.Song) *exception.InternalServerError {
	if err := s.repository.Delete(song); err != nil {
		return exception.NewInternalServerError("Gagal menghapus lagu")
	}

	return nil
}
