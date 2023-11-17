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
	GinReadTimeout  int
	GinWriteTimeout int
	GinIdleTimeout	int
	// Redis credentials
	RedisURI	string
	RedisPass	string
	RedisDB		int
	// JWT Token credentials
	AccessTokenPrivateKey	string
	AccessTokenPublicKey	string
	AccessTokenExpiresIn	int
	AccessTokenMaxAge		int
	RefreshTokenPrivateKey	string
	RefreshTokenPublicKey	string
	RefreshTokenExpiresIn	int
	RefreshTokenMaxAge		int
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

	accTknPvtKey := os.Getenv("ACCESS_TKN_PRV_KEY")
	accTknPubKey :=os.Getenv("ACCESS_TKN_Pub_KEY")
	accTknExpiresIn := os.Getenv("ACCESS_TKN_EXP_IN")
	accTknExpIn, err := strconv.Atoi(accTknExpiresIn)
	if err != nil {
		fmt.Println(err)	}
	accTknMaxage := os.Getenv("ACESS_TKN_MAXAGE")
	accTknMaxAge, err := strconv.Atoi(accTknMaxage)
	if err != nil {
		fmt.Println(err)	}
	refTknPvtKey := os.Getenv("REFRESH_TKN_PRV_KEY")
	refTknPubKey :=os.Getenv("REFRESH_TKN_PUB_KEY")
	refTknExpiresIn := os.Getenv("REFRESH_TKN_EXP_IN")
	refTknExpIn, err := strconv.Atoi(refTknExpiresIn)
	if err != nil {
		fmt.Println(err)	}
	refTknMaxage := os.Getenv("REFRESH_TKN_MAXAGE")
	refTknMaxAge, err := strconv.Atoi(refTknMaxage)
	if err != nil {
		fmt.Println(err)	}

	log.Println("[INIT] ✅ configuration loaded")
	
	return &Config{
		PSQL_DSN:        dsn,
		GinPort:         ginPort,
		GinReadTimeout:  gRTO,
		GinWriteTimeout: gWTO,
		GinIdleTimeout:  gITO,
		RedisURI:        redisuri,
		RedisPass:       redispwd,
		RedisDB:         redis_db,
		AccessTokenPrivateKey: accTknPvtKey,
		AccessTokenPublicKey: accTknPubKey,
		AccessTokenExpiresIn: accTknExpIn,
		AccessTokenMaxAge: accTknMaxAge,
		RefreshTokenPrivateKey: refTknPvtKey,
		RefreshTokenPublicKey: refTknPubKey,
		RefreshTokenExpiresIn: refTknExpIn,
		RefreshTokenMaxAge: refTknMaxAge,
	}
}
