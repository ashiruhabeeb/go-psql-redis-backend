package main

import (
	"log"

	"github.com/ashiruhabeeb/go-backend/db"
	"github.com/ashiruhabeeb/go-backend/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadAppConfig()

	db.PostgresConnect(cfg.PSQL_DSN)

	db.RedisConnect(cfg)

	gn := gin.Default()

	log.Printf("[INIT] ✅ gin router running and listening on port %v", cfg.GinPort)
	
	gn.Run(":"+cfg.GinPort)
}
