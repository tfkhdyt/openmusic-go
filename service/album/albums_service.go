package album

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	albumRepository "github.com/tfkhdyt/openmusic-go/repository/postgres/album"
)

type Service struct {
	repository *albumRepository.Repository
}

func NewService(repository *albumRepository.Repository) *Service {
	return &Service{repository}
}

func (s Service) AddAlbum(name string, year uint16) (gin.H, error) {
	id := fmt.Sprintf("album-%v", uuid.NewString())

	album, err := s.repository.CreateAlbum(id, name, year)

	data := gin.H{
		"albumId": album.ID,
	}

	return data, err
}
