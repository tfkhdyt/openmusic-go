package collab

import "github.com/tfkhdyt/openmusic-go/repository/postgres/collab"

type Service struct {
	repository *collab.Repository
}

func NewService(repo *collab.Repository) *Service {
	return &Service{repo}
}
