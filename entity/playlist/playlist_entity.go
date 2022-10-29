package playlist

type Playlist struct {
	ID     string `json:"id"`
	Name   string `json:"name" binding:"required"`
	UserId string `json:"owner"`
}
