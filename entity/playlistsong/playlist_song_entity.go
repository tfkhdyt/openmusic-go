package playlistsong

type PlaylistSong struct {
	SongId string `json:"songId" binding:"required"`
}
