package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GinPort	string
}

func LoadAppConfig() *Config {
	godotenv.Load()


	srvPort := os.Getenv("GIN_PORT")

	return &Config{
		GinPort: srvPort,
	}
}
