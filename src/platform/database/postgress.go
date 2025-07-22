package database

import (
	"database/sql"
	"log"
	"sync"
)

var postgressDb *sql.DB
var once sync.Once

func NewPostgressDb(dsn string) (*sql.DB, error) {
	var err error
	once.Do(func() {
		postgressDb, err = sql.Open("postgress", dsn)
	})

	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return nil, err
	}

	return postgressDb, nil
}
