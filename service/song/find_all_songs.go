package song

import (
	"fmt"

	"github.com/tfkhdyt/openmusic-go/entity/song"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) FindAll(title string, performer string) ([]song.Song, *exception.InternalServerError) {
	if title != "" {
		title = fmt.Sprintf("%%%v%%", title)
	}
	if performer != "" {
		performer = fmt.Sprintf("%%%v%%", performer)
	}

	fmt.Println(title, performer)

	songs, err := s.repository.FindAll(title, performer)
	if err != nil {
		return songs, exception.NewInternalServerError("Gagal mendapatkan lagu")
	}

	return songs, nil
}
