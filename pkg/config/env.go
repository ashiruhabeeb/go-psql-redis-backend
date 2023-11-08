package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Cfg *Config
type Config struct {
	// Postgres credentials
	PSQL_DSN	string
	// Gin router credentials
	GinPort		string
	// Redis credentials
	RedisURI	string
	RedisPass	string
	RedisDB		int
}

func LoadAppConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ERROR] godotenv.Load failure :%v", err)
	}

	dsn := os.Getenv("DB_DSN")
	ginPort := os.Getenv("GIN_PORT")
	redisuri := os.Getenv("REDIS_URL")
	redispwd := os.Getenv("REDIS_PASSWORD")
	redisdb := os.Getenv("REDIS_DB")
	redis_db, err := strconv.Atoi(redisdb)
	if err != nil {
		fmt.Println(err)
	}

	return &Config{
		PSQL_DSN:  dsn,
		GinPort:   ginPort,
		RedisURI:  redisuri,
		RedisPass: redispwd,
		RedisDB:   redis_db,
	}
}
