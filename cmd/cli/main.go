package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/service"
	"github.com/jxcf/jxcf-api/internal/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Validate required environment variables
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: DEEPSEEK_API_KEY is not set")
		os.Exit(1)
	}

	// Load configuration
	cfg := config.Load()

	// Example article for testing
	article := `Artificial intelligence is transforming modern business. Companies worldwide are adopting AI to enhance their operations. Machine learning models are becoming increasingly sophisticated and accessible.`

	fmt.Println("=== jxcf-api CLI Demo ===")
	fmt.Println()
	fmt.Println("Article:")
	fmt.Println(article)
	fmt.Println()

	// Create analyzer
	analyzer := service.NewAnalyzer(cfg)

	// Analyze article
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("Analyzing article...")
	resp, err := analyzer.Analyze(ctx, models.AnalyzeRequest{
		Article:  article,
		Language: "en",
	})

	if err != nil {
		log.Fatalf("Analysis failed: %v", err)
	}

	// Display results
	fmt.Println("\nResults:")
	fmt.Println("--------")
	fmt.Printf("Summary: %s\n", resp.Summary)
	fmt.Printf("Meta Description: %s\n", resp.MetaDescription)
	fmt.Printf("SEO Score: %.1f/10\n", resp.SEOScore)
	fmt.Println("\nKeywords:")
	for i, kw := range resp.Keywords {
		fmt.Printf("  %d. %s (Relevance: %.2f)\n", i+1, kw.Keyword, kw.Relevance)
	}
}
