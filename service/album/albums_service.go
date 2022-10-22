package album

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tfkhdyt/openmusic-go/repository/postgres/album"
)

type Service struct {
	repository *album.Repository
}

func NewService(repository *album.Repository) *Service {
	return &Service{repository}
}

func (s Service) AddAlbum(name string, year uint16) {
	id := fmt.Sprintf("album-%v", uuid.NewString())

	fmt.Println("ID:", id)
}
