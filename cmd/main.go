package main

import (
	"github.com/ashiruhabeeb/go-backend/app/router"
	"github.com/ashiruhabeeb/go-backend/db"
	"github.com/ashiruhabeeb/go-backend/pkg/config"
)

func main() {
	cfg := config.LoadAppConfig()

	db.PostgresConnect(cfg.PSQL_DSN)

	db.RedisConnect(cfg)

	router.SetupGinRouter(cfg)
}
