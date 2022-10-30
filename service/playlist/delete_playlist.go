package playlist

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Delete(playlist *playlist.Playlist) *exception.InternalServerError {
	if err := s.repository.Delete(playlist); err != nil {
		return exception.NewInternalServerError("Gagal menghapus playlist")
	}

	return nil
}
