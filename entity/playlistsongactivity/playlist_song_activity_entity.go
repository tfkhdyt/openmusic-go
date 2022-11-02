package playlistsongactivity

import "time"

type PlaylistSongActivity struct {
	ID         string
	PlaylistID string    `json:"playlistId"`
	SongID     string    `json:"songId"`
	UserID     string    `json:"userId"`
	Action     string    `json:"action"`
	Time       time.Time `json:"time"`
}

type FindAllActivities struct {
	Username string    `json:"username"`
	Title    string    `json:"title"`
	Action   string    `json:"action"`
	Time     time.Time `json:"time"`
}
