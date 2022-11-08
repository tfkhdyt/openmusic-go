package playlist

import (
	"github.com/tfkhdyt/openmusic-go/repository/postgres/playlist"
	"github.com/tfkhdyt/openmusic-go/service/collab"
)

type Service struct {
	repository     *playlist.Repository
	collabsService *collab.Service
}

func NewService(repository *playlist.Repository, collabsService *collab.Service) *Service {
	return &Service{repository, collabsService}
}
