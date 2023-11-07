package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Cfg *Config
type Config struct {
	// Postgres credentials
	PSQL_DSN	string
	// Postgres Credentials for test
	PSQL_DSNN	string
	// Gin router credentials
	GinPort		string
	// Redis credentials
	RedisURI	string
	RedisPass	string
	RedisDB		int

}

func LoadAppConfig() *Config {
	godotenv.Load()

	dsn := os.Getenv("DB_DSN")
	dsnn := os.Getenv("DB_DSNN")
	ginPort := os.Getenv("GIN_PORT")
	redisuri := os.Getenv("REDIS_URL")
	redispwd := os.Getenv("REDIS_PASSWORD")
	redisdb := os.Getenv("REDIS_DB")
	redis_db, err := strconv.Atoi(redisdb)
	if err != nil {
		panic(err)
	}

	return &Config{
		PSQL_DSN:  dsn,
		PSQL_DSNN: dsnn,
		GinPort:   ginPort,
		RedisURI:  redisuri,
		RedisPass: redispwd,
		RedisDB:   redis_db,
	}
}
