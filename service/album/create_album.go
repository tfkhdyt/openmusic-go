package album

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s *Service) Create(name string, year uint16) (string, *exception.InternalServerError) {
	id := fmt.Sprintf("album-%v", uuid.NewString())

	album, err := s.repository.Create(id, name, year)
	if err != nil {
		return "", exception.NewInternalServerError("Gagal menambahkan album")
	}

	return album.ID, nil
}
