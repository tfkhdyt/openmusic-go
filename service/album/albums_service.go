package album

import (
	"fmt"

	"github.com/google/uuid"
	albumEntity "github.com/tfkhdyt/openmusic-go/entity/album"
	"github.com/tfkhdyt/openmusic-go/exception"
	albumRepository "github.com/tfkhdyt/openmusic-go/repository/postgres/album"
)

type Service struct {
	repository *albumRepository.Repository
}

func NewService(repository *albumRepository.Repository) *Service {
	return &Service{repository}
}

func (s Service) Create(name string, year uint16) (string, *exception.InternalServerError) {
	id := fmt.Sprintf("album-%v", uuid.NewString())

	album, err := s.repository.Create(id, name, year)
	if err != nil {
		return "", exception.NewInternalServerError("Gagal menambahkan album")
	}

	return album.ID, nil
}

func (s Service) FindOne(id string) (albumEntity.Album, *exception.NotFoundError) {
	album, err := s.repository.FindOne(id)
	if err != nil {
		return album, exception.NewNotFoundError("Album tidak ditemukan")
	}

	return album, nil
}

func (s Service) Update(oldAlbum *albumEntity.Album, newAlbum *albumEntity.Album) *exception.InternalServerError {
	if err := s.repository.Update(oldAlbum, newAlbum); err != nil {
		return exception.NewInternalServerError("Gagal mengubah album")
	}

	return nil
}

func (s Service) Delete(album *albumEntity.Album) *exception.InternalServerError {
	if err := s.repository.Delete(album); err != nil {
		return exception.NewInternalServerError("Gagal menghapus album")
	}

	return nil
}
