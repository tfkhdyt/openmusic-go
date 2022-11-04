package collab

type Collaboration struct {
	PlaylistId string `json:"playlistId" binding:"required"`
	UserId     string `json:"userId" binding:"required"`
}
