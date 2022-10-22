package album

type Album struct {
	ID   string `gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Year uint16 `json:"year" binding:"required"`
}
