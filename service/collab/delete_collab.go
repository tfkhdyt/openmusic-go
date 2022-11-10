package collab

import (
	"github.com/tfkhdyt/openmusic-go/entity/collab"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Delete(collab *collab.Collaboration) *exception.InternalServerError {
	if err := s.repository.Delete(collab); err != nil {
		return exception.NewInternalServerError("Gagal menghapus kolaborasi")
	}

	return nil
}
