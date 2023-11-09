package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func PostgresConnect(dsn string) *sql.DB {
	// Establish PSQL connection based on parameter provided
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("[ERROR] database connection failed: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	log.Println("[INIT] ✅ postgresql database connection established")

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(5)

	return db
}
