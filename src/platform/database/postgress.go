package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewPostgresDB() *sql.DB {
	dsn := viper.GetString("POSTGRES_DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open Postgres: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping Postgres: %v", err)
	}
	log.Println("Connected to Postgres")
	return db
}
