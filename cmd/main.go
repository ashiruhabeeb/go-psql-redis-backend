package main

import (
	"github.com/ashiruhabeeb/go-backend/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadAppConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	gn := gin.Default()

	gn.Run(":"+cfg.GinPort)
}