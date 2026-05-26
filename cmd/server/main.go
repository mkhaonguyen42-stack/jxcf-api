package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/handlers"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Validate required environment variables
	if os.Getenv("DEEPSEEK_API_KEY") == "" {
		fmt.Println("Error: DEEPSEEK_API_KEY is not set")
		os.Exit(1)
	}

	// Load configuration
	cfg := config.Load()

	// Start server
	if err := handlers.StartServer(cfg); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
