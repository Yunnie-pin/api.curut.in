package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env            string
	Port           string
	SecretKey      string
	TokenExpiresIn string
	TokenMaxAge    string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBURL          string
}

func LoadConfig() (config Config, err error) {
	godotenv.Load()

	config = Config{
		Env:            os.Getenv("ENV"),
		Port:           os.Getenv("PORT"),
		SecretKey:      os.Getenv("SECRET"),
		TokenExpiresIn: os.Getenv("TOKEN_EXPIRED_IN"),
		TokenMaxAge:    os.Getenv("TOKEN_MAXAGE"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBURL:          os.Getenv("DB_URL"),
	}

	return
}
