package auth

import "github.com/tfkhdyt/openmusic-go/entity/auth"

func (r Repository) FindOne(token string) error {
	var auth auth.Auth

	if err := r.db.First(&auth, "token = ?", token).Error; err != nil {
		return err
	}

	return nil
}
