package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sbdoddi7/innoscripta/src/platform/database"
	logger "github.com/sbdoddi7/innoscripta/src/platform/log"
	"github.com/sbdoddi7/innoscripta/src/platform/queue"
	"github.com/sbdoddi7/innoscripta/src/routes"
)

func main() {
	// Initialize logger
	logger.Init()
	logger.Logger.Info("Innoscripta bank app starting...")

	// Connect to Postgres
	postgresDB, err := database.NewPostgresDB()
	if err != nil {
		logger.Logger.Fatalf("Failed to connect to Postgres DB: %v", err)
	}
	defer postgresDB.Close()

	// Connect to MongoDB
	mongoClient, err := database.NewMongoClient()
	if err != nil {
		logger.Logger.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			logger.Logger.Errorf("Error disconnecting MongoDB: %v", err)
		}
	}()

	// Connect to RabbitMQ
	rabbitCh, rabbitConn, err := queue.NewRabbitMQChannel()
	if err != nil {
		logger.Logger.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitCh.Close()
	defer rabbitConn.Close()

	// Setup router with dependencies
	router := routes.NewRouter(postgresDB, mongoClient, rabbitCh)

	// Read port from env (default 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Logger.Infof("Server will start at :%s", port)

	// Channel to catch OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		if err := router.Run(":" + port); err != nil {
			logger.Logger.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for quit signal
	<-quit
	logger.Logger.Warn("Shutting down gracefully...")

	// Optional: sleep or cleanup
	time.Sleep(2 * time.Second)

	logger.Logger.Info("Server stopped")
}
