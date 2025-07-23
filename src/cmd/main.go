package main

import (
	"os"

	logger "github.com/sbdoddi7/innoscripta/src/platform/log"
	"github.com/sbdoddi7/innoscripta/src/routes"
)

func main() {
	logger.Init()
	logger.Logger.Info("Innoscripta bank app starting...")

	router := routes.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Logger.Infof("Server started at :%s", port)

	if err := router.Run(":" + port); err != nil {
		logger.Logger.Fatal(err)
	}
}
