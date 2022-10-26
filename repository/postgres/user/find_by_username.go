package user

import "github.com/tfkhdyt/openmusic-go/entity/user"

func (r Repository) FindByUsername(username string) (user.User, error) {
	var user user.User

	if err := r.db.Select("id", "password").Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
