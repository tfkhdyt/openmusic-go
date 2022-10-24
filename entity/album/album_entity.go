package album

import "github.com/tfkhdyt/openmusic-go/entity/song"

type Album struct {
	ID    string      `json:"id"`
	Name  string      `json:"name" binding:"required"`
	Year  uint16      `json:"year" binding:"required"`
	Songs []song.Song `json:"songs,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
