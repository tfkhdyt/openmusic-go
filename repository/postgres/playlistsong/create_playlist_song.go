package playlistsong

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/entity/song"
)

func (r Repository) Create(playlist *playlist.Playlist, song *song.Song) error {
	if err := r.db.Model(playlist).Association("Songs").Append(song); err != nil {
		return err
	}

	return nil
}
