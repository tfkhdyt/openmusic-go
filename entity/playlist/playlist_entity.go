package playlist

import "github.com/tfkhdyt/openmusic-go/entity/song"

type Playlist struct {
	ID     string      `json:"id"`
	Name   string      `json:"name" binding:"required"`
	UserId string      `json:"owner"`
	Songs  []song.Song `json:"songs" gorm:"many2many:playlist_songs;"`
}
