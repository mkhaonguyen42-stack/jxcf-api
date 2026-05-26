package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/llm"
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
	article := `人工智能上次性热潮是年一年成的翿fn货。下面来看人工智能的上一波翿fn货了。一个新板上架了，就能接统每天的首页能流量。这个新板上架了，到底有什么不一样？
`

	fmt.Println("=== jxcf-api CLI Demo ===")
	fmt.Println()
	fmt.Println("Article:")
	fmt.Println(article)
	fmt.Println()

	// Create client and services
	client := llm.NewDeepSeekClient(
		cfg.DeepseekAPIKey,
		cfg.DeepseekBaseURL,
		cfg.DeepseekModel,
	)

	analyzer := service.NewAnalyzer(cfg)

	// Analyze article
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("Analyzing article...")
	resp, err := analyzer.Analyze(ctx, models.AnalyzeRequest{
		Article: article,
		Language: "zh",
	})

	if err != nil {
		log.Fatalf("Analysis failed: %v", err)
	}

	// Display results
	fmt.Println("Results:")
	fmt.Println("--------")
	fmt.Printf("Summary: %s\n", resp.Summary)
	fmt.Printf("Meta Description: %s\n", resp.MetaDescription)
	fmt.Printf("SEO Score: %.1f/10\n", resp.SEOScore)
	fmt.Println()
	fmt.Println("Keywords:")
	for i, kw := range resp.Keywords {
		fmt.Printf("  %d. %s (%.2f)\n", i+1, kw.Keyword, kw.Relevance)
	}
}
`

// Let me fix the import
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/llm"
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
	article := `人工智能上次性热潮是年一年成的ффн货。下面来看人工智能的上一波ффн货了。一个新板上架了，就能接统每天的首页能流量。这个新板上架了，到底有什么不一样？`

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
		Article: article,
		Language: "zh",
	})

	if err != nil {
		log.Fatalf("Analysis failed: %v", err)
	}

	// Display results
	fmt.Println("Results:")
	fmt.Println("--------")
	fmt.Printf("Summary: %s\n", resp.Summary)
	fmt.Printf("Meta Description: %s\n", resp.MetaDescription)
	fmt.Printf("SEO Score: %.1f/10\n", resp.SEOScore)
	fmt.Println()
	fmt.Println("Keywords:")
	for i, kw := range resp.Keywords {
		fmt.Printf("  %d. %s (%.2f)\n", i+1, kw.Keyword, kw.Relevance)
	}
}
`
	
Let me re-do this properly: