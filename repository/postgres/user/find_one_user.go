package user

import "github.com/tfkhdyt/openmusic-go/entity/user"

func (r Repository) FindOne(id string) (user.User, error) {
	var user user.User

	err := r.db.First(&user, "id = ?", id).Error

	return user, err
}
