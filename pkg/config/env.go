package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PSQL_DSN	string
	GinPort		string
}

func LoadAppConfig() *Config {
	godotenv.Load()


	srvPort := os.Getenv("GIN_PORT")

	return &Config{
		GinPort: srvPort,
	}
}
