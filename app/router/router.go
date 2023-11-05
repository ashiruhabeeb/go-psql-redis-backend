package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupGinRouter(g *gin.Engine) {
	g.Use(gin.Logger())

	log.Println("[INIT] ✅ gin router set up")

	g.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
