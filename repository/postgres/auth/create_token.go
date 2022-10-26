package auth

import "github.com/tfkhdyt/openmusic-go/entity/auth"

func (r Repository) Store(token *auth.Auth) error {
	err := r.db.Create(token).Error

	return err
}
