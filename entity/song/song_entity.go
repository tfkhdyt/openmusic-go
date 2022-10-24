package song

type Song struct {
	ID        string `gorm:"primaryKey"`
	Title     string `json:"title" binding:"required"`
	Year      uint16 `json:"year" binding:"required"`
	Performer string `json:"performer" binding:"required"`
	Genre     string `json:"genre" binding:"required"`
	Duration  uint   `json:"duration" binding:"required"`
	AlbumID   string `json:"albumId" binding:"required"`
}
