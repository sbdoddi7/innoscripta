package main

import (
	"log"
	"os"

	"github.com/sbdoddi7/innoscripta/src/routes"
)

func main() {
	// init DB, queue, envs etc (skipped)

	router := routes.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started at :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
