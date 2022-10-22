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

func (s Service) Add(name string, year uint16) (string, error) {
	id := fmt.Sprintf("album-%v", uuid.NewString())

	album, err := s.repository.Add(id, name, year)
	if err != nil {
		return "", errors.New("gagal menambahkan album")
	}

	return album.ID, nil
}

func (s Service) GetById(id string) (albumEntity.Album, error) {
	album, err := s.repository.GetById(id)
	if err != nil {
		return album, errors.New("album tidak ditemukan")
	}

	return album, nil
}
