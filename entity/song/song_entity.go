package song

type Song struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Title     string `json:"title,omitempty" binding:"required"`
	Year      uint16 `json:"year,omitempty" binding:"required"`
	Performer string `json:"performer,omitempty" binding:"required"`
	Genre     string `json:"genre,omitempty" binding:"required"`
	Duration  uint   `json:"duration,omitempty"`
	AlbumID   string `json:"albumId,omitempty"`
}
