package playlistsong

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Delete(playlist *playlist.Playlist, song *song.Song) *exception.InternalServerError {
	if err := s.repository.Delete(playlist, song); err != nil {
		return exception.NewInternalServerError("Gagal menghapus lagu pada playlist")
	}

	return nil
}
