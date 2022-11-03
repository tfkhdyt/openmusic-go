package playlist

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlistsongactivity"
	"github.com/tfkhdyt/openmusic-go/entity/song"
)

type Playlist struct {
	ID         string                                      `json:"id"`
	Name       string                                      `json:"name" binding:"required"`
	UserId     string                                      `json:"owner,omitempty"`
	Songs      []song.Song                                 `json:"songs,omitempty" gorm:"many2many:playlist_songs"`
	Activities []playlistsongactivity.PlaylistSongActivity `json:"activities,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	// Users      []user.User                                 `json:"users" gorm:"many2many:user_playlists"`
}

type FindAllPlaylistsResult struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
