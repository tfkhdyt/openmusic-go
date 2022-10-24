package song

import "github.com/tfkhdyt/openmusic-go/repository/postgres/song"

type Service struct {
	repository *song.Repository
}

func NewService(repository *song.Repository) *Service {
	return &Service{repository}
}
