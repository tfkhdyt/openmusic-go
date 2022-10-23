package album

import (
	albumsRepository "github.com/tfkhdyt/openmusic-go/repository/postgres/album"
)

type Service struct {
	repository *albumsRepository.Repository
}

func NewService(repository *albumsRepository.Repository) *Service {
	return &Service{repository}
}
