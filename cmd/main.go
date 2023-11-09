package main

import (
	"github.com/ashiruhabeeb/go-backend/app/router"
	"github.com/ashiruhabeeb/go-backend/db"
	"github.com/ashiruhabeeb/go-backend/pkg/config"
)

func main() {
	// Load app environment variables
	cfg := config.LoadAppConfig()
	// Initialize PSQL db connection 
	psql := db.PostgresConnect(cfg.PSQL_DSN)
	// Initialize Redis db connection
	db.RedisConnect(cfg.RedisURI, cfg.RedisPass, cfg.RedisDB)
	// Initialize app roouter
	router.SetupGinRouter(
		psql,
		cfg.GinPort,
		cfg.GinReadTimeout,
		cfg.GinWriteTimeout,
		cfg.GinIdleTimeout,
	)
}
