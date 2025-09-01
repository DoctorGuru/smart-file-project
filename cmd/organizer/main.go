package main

import (
	"log"
	"smart-file-organizer/internal/app"
	"smart-file-organizer/internal/transport"
	"smart-file-organizer/pkg/config"
	"smart-file-organizer/pkg/logger"
)

func main() {

	rootCmd := transport.NewCLI()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	// Load config
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Init logger
	logFile := logger.Init("organizer.log")
	defer logFile.Close()

	// Run organizer
	app.Run(cfg)
}
