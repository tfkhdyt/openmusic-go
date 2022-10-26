package token

import (
	"github.com/tfkhdyt/openmusic-go/config/jwt"
)

type Manager struct {
	config *jwt.Config
}

func NewManager(config *jwt.Config) *Manager {
	return &Manager{config}
}
