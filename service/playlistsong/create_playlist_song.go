package playlistsong

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Create(playlist *playlist.Playlist, song *song.Song) *exception.InternalServerError {
	if err := s.repository.Create(playlist, song); err != nil {
		return exception.NewInternalServerError("Gagal menambahkan lagu ke dalam playlsit")
	}

	return nil
}
