package playlist

import "github.com/tfkhdyt/openmusic-go/repository/postgres/playlist"

type Service struct {
	repository *playlist.Repository
}

func NewService(repository *playlist.Repository) *Service {
	return &Service{repository}
}
