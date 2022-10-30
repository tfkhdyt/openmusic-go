package playlistsong

import "github.com/tfkhdyt/openmusic-go/repository/postgres/playlistsong"

type Service struct {
	repository *playlistsong.Repository
}

func NewService(repository *playlistsong.Repository) *Service {
	return &Service{repository}
}
