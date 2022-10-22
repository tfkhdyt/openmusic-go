package postgres

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	DBName   string
	Port     uint16
	User     string
	Password string
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
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

func (c Config) GetPassword() string {
	return os.Getenv("DB_PASS")
}
