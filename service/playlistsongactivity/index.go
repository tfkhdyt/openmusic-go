package playlistsongactivity

import "github.com/tfkhdyt/openmusic-go/repository/postgres/playlistsongactivity"

type Service struct {
	repository *playlistsongactivity.Repository
}

func NewService(repo *playlistsongactivity.Repository) *Service {
	return &Service{repo}
}
