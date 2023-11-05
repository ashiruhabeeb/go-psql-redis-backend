package main

import (
	"github.com/ashiruhabeeb/go-backend/db"
	"github.com/ashiruhabeeb/go-backend/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadAppConfig()

	db.PostgresConnect(cfg.PSQL_DSN)

	gn := gin.Default()

	gn.Run(":"+cfg.GinPort)
}
