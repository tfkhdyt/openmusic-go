package user

import (
	"github.com/tfkhdyt/openmusic-go/entity/auth"
	"github.com/tfkhdyt/openmusic-go/entity/user"
)

func (r Repository) FindByCredentials(credentials *auth.Credentials) (string, error) {
	var user user.User

	err := r.db.Where(credentials).First(&user).Error

	return user.ID, err
}
