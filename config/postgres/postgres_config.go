package postgres

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		// panic("Error loading .env file")
		log.Println(err.Error())
	}
	return &Config{}
}

func (c Config) GetHost() string {
	return os.Getenv("DB_HOST")
}

func (c Config) GetDBName() string {
	return os.Getenv("DB_NAME")
}

func (c Config) GetPort() string {
	return os.Getenv("DB_PORT")
}

func (c Config) GetUser() string {
	return os.Getenv("DB_USER")
}

func (c Config) GetPass() string {
	return os.Getenv("DB_PASS")
}
