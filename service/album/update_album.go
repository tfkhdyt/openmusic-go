package album

import (
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s *Service) Update(oldAlbum *albumEntity.Album, newAlbum *albumEntity.Album) *exception.InternalServerError {
	if err := s.repository.Update(oldAlbum, newAlbum); err != nil {
		return exception.NewInternalServerError("Gagal mengubah album")
	}

	return nil
}
