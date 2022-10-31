package playlistsong

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/entity/song"
)

func (r Repository) Delete(playlist *playlist.Playlist, song *song.Song) error {
	if err := r.db.Model(playlist).Association("Songs").Delete(song); err != nil {
		return err
	}

	return nil
}
