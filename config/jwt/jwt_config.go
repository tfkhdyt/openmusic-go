package jwt

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}
	return &Config{}
}

func (c Config) GetSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY")
}
