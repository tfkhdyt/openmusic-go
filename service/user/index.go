package user

import "github.com/tfkhdyt/openmusic-go/repository/postgres/user"

type Service struct {
	repository *user.Repository
}

func NewService(repository *user.Repository) *Service {
	return &Service{repository}
}
