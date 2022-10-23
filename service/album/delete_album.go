package album

import (
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s *Service) Delete(album *albumEntity.Album) *exception.InternalServerError {
	if err := s.repository.Delete(album); err != nil {
		return exception.NewInternalServerError("Gagal menghapus album")
	}

	return nil
}
