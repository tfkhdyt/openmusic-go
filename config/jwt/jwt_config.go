package jwt

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	return &Config{}
}

func (c Config) GetSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY")
}
