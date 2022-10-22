package album

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
	albumRepository "github.com/tfkhdyt/openmusic-go/repository/postgres/album"
)

type Service struct {
	repository *albumRepository.Repository
}

func NewService(repository *albumRepository.Repository) *Service {
	return &Service{repository}
}

func (s Service) Create(name string, year uint16) (string, error) {
	id := fmt.Sprintf("album-%v", uuid.NewString())

	album, err := s.repository.Create(id, name, year)
	if err != nil {
		return "", errors.New("gagal menambahkan album")
	}

	return album.ID, nil
}

func (s Service) FindOne(id string) (albumEntity.Album, error) {
	album, err := s.repository.FindOne(id)
	if err != nil {
		return album, errors.New("album tidak ditemukan")
	}

	return album, nil
}
