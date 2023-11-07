package main_test

import (
	"log"
	"os"
	"testing"

	"github.com/ashiruhabeeb/go-backend/app/router"
	"github.com/ashiruhabeeb/go-backend/db"
	"github.com/ashiruhabeeb/go-backend/pkg/config"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	cfg := config.LoadAppConfig()

	db.PostgresConnect(cfg.PSQL_DSNN)

	db.RedisConnect(cfg)

	g := gin.Default()

	log.Printf("[INIT] ✅ gin router running and listening on port %v", cfg.GinPort)

	router.SetupGinRouter(g)

	code := m.Run()
	os.Exit(code)
}
