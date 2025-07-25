package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	dsn := os.Getenv("POSTGRES_DSN")
	var db *sql.DB
	var err error

	for range 10 {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Connected to Postgres!")
				return db, nil
			}
		}

		log.Printf("Postgres not ready yet (%v). Retrying in 3s...", err)
		// back-off retry after 3 seconds
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to Postgres after retries: %w", err)
}
