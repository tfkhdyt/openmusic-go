package album

import (
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s *Service) FindOne(id string) (albumEntity.Album, *exception.NotFoundError) {
	album, err := s.repository.FindOne(id)
	if err != nil {
		return album, exception.NewNotFoundError("Album tidak ditemukan")
	}

	return album, nil
}
