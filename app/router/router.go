package router

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ashiruhabeeb/go-backend/handlers"
	"github.com/ashiruhabeeb/go-backend/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupGinRouter(db *sql.DB, port string, r, w, i int) {
	gn := gin.Default()
	gn.Use(gin.Logger())

	gn.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return true
        },
        MaxAge: 15 * time.Second,
	}))

	gn.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	storageRepo := storage.NewUserStorage(db)
	usersHandlers := handlers.NewUsersHandlers(storageRepo)

	users := gn.Group("/api/v1")
	users.POST("/signup", usersHandlers.UserSignUp)
	users.GET("/user/:id", usersHandlers.GetUserById)
	users.GET("/fetch/:email", usersHandlers.GetUserByEmail)
	users.GET("/get/:username", usersHandlers.GetUserByUsername)
	users.GET("/allrecords", usersHandlers.FetchAllUsersRecords)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      gn,
		MaxHeaderBytes: 1 << 20, //1 MB
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
	
	// declare a buffered channel that reveives unix signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	log.Println("[CLOSE] shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("[CLOSE] error when shutdown server: %s", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("[CLOSE] timeout of 5 seconds.")
	log.Println("[CLOSE] server exiting")
}
