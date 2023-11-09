package config

import (
	"fmt"
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
	GinReadTimeout  int
	GinWriteTimeout int
	GinIdleTimeout	int
	// Redis credentials
	RedisURI	string
	RedisPass	string
	RedisDB		int
}

func LoadAppConfig() *Config {
	godotenv.Load()
	
	dsn := os.Getenv("DB_DSN")

	ginPort := os.Getenv("GIN_PORT")
	ginReadTO := os.Getenv("GIN_READTIMEOUT")
	gRTO, err := strconv.Atoi(ginReadTO)
	if err != nil {
		fmt.Println(err)
	}
	ginWriteTO := os.Getenv("GIN_WRITETIMEOUT")
	gWTO, err := strconv.Atoi(ginWriteTO)
	if err != nil {
		fmt.Println(err)
	}
	ginIdleTO := os.Getenv("GIN_IDLETIMEOUT")
	gITO, err := strconv.Atoi(ginIdleTO)
	if err != nil {
		fmt.Println(err)
	}
	
	redisuri := os.Getenv("REDIS_URL")
	redispwd := os.Getenv("REDIS_PASSWORD")
	redisdb := os.Getenv("REDIS_DB")
	redis_db, err := strconv.Atoi(redisdb)
	if err != nil {
		fmt.Println(err)
	}

	return &Config{
		PSQL_DSN:        dsn,
		GinPort:         ginPort,
		GinReadTimeout:  gRTO,
		GinWriteTimeout: gWTO,
		GinIdleTimeout:  gITO,
		RedisURI:        redisuri,
		RedisPass:       redispwd,
		RedisDB:         redis_db,
	}
}
