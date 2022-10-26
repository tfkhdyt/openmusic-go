package auth

import "github.com/tfkhdyt/openmusic-go/repository/postgres/auth"

type Service struct {
	repository *auth.Repository
}

func NewService(repository *auth.Repository) *Service {
	return &Service{repository}
}
