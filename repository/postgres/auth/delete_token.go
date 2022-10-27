package auth

import "github.com/tfkhdyt/openmusic-go/entity/auth"

func (r Repository) DeleteToken(token string) error {
	err := r.db.Delete(&auth.Auth{}, "token = ?", token).Error

	return err
}
