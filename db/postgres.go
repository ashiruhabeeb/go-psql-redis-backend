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

	err = db.Ping()
	if err != nil {
		log.Printf("[ERROR] db.Ping failure: %s", err)
	}

	log.Println("[INIT] ✅ postgresql database connection established")

	// 
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5*time.Minute)

	return db
}
