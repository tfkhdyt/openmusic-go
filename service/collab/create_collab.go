package collab

import (
	"github.com/tfkhdyt/openmusic-go/entity/playlist"
	"github.com/tfkhdyt/openmusic-go/entity/user"
	"github.com/tfkhdyt/openmusic-go/exception"
)

func (s Service) Create(playlist *playlist.Playlist, user *user.User) *exception.InternalServerError {
	if err := s.repository.Create(playlist, user); err != nil {
		return exception.NewInternalServerError("Gagal menambahkan collab")
	}

	return nil
}
