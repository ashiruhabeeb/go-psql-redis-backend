package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupGinRouter(port string, r, w, i int) {
	gn := gin.Default()
	gn.Use(gin.Logger())

	gn.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      gn,
		ReadTimeout:  time.Duration(time.Duration(r).Seconds()),
		WriteTimeout: time.Duration(time.Duration(w).Seconds()),
		IdleTimeout:  time.Duration(time.Duration(i).Seconds()),
	}

	go func(){
		log.Printf("[INIT] ✅ gin router running and listening on port %v", port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] http.ListenAndServe failure: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("[ERROR] Server Shutdown failure:", err)
	}
	
	log.Println("Server exiting")
}
