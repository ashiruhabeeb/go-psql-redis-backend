package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func PostgresConnect(dsn string) *sql.DB {
	// Establish PSQL connection based on parameter provided
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("[ERROR] database connection failed: %v", err)
		os.Exit(1)
	}

	// Verify established database connection status
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		log.Println("[INIT] ✅ postgresql database pinged")
	}

	log.Println("[INIT] ✅ postgresql database connection established")

	// Database connection management settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5*time.Minute)

	return db
}
