package album

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tfkhdyt/openmusic-go/entity/album"
	albumRepository "github.com/tfkhdyt/openmusic-go/repository/postgres/album"
)

type Service struct {
	repository *albumRepository.Repository
}

func NewService(repository *albumRepository.Repository) *Service {
	return &Service{repository}
}

func (s Service) Add(name string, year uint16) (gin.H, error) {
	id := fmt.Sprintf("album-%v", uuid.NewString())

	album, err := s.repository.Add(id, name, year)
	if err != nil {
		return gin.H{}, errors.New("gagal menambahkan album")
	}

	data := gin.H{
		"albumId": album.ID,
	}

	return data, nil
}

func (s Service) GetById(id string) (album.Album, error) {
	album, err := s.repository.GetById(id)
	if err != nil {
		return album, errors.New("album tidak ditemukan")
	}

	return album, nil
}
