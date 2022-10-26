package user

import "github.com/tfkhdyt/openmusic-go/entity/user"

func (r Repository) Create(user *user.User) (string, error) {
	err := r.db.Create(user).Error

	return user.ID, err
}
