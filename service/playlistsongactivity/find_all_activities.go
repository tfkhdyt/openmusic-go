package playlistsongactivity

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlistsongactivity"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) FindAll(id string) ([]playlistsongactivity.FindAllActivities, *exception.NotFoundError) {
	activities, err := s.repository.FindAll(id)
	if err != nil {
		return activities, exception.NewNotFoundError("Aktivitas tidak ditemukan")
	}

	return activities, nil
}
