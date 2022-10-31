package playlist

import "github.com/tfkhdyt/openmusic-go/entity/song"

type Playlist struct {
	ID     string      `json:"id"`
	Name   string      `json:"name" binding:"required"`
	UserId string      `json:"owner,omitempty"`
	Songs  []song.Song `json:"songs,omitempty" gorm:"many2many:playlist_songs;constraint:OnDelete:CASCADE"`
}
