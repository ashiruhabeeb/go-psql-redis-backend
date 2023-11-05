package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func PostgresConnect(dsn string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("[ERROR] database connection failed: %v", err)
	}
	defer db.Close()

	log.Println("[INIT] ✅ postgresql database connection established")

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(5)
}
