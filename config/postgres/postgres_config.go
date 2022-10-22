package postgres

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	DbName   string
	Port     uint16
	Username string
	Password string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Failed to parse port to integer")
	}

	return &Config{
		Host:     os.Getenv("DB_HOST"),
		DbName:   os.Getenv("DB_NAME"),
		Port:     uint16(port),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}
}
