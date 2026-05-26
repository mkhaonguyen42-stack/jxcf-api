package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/models"
)

type BatchAnalyzer struct {
	analyzer *Analyzer
	cfg *config.Config
}

func NewBatchAnalyzer(cfg *config.Config) *BatchAnalyzer {
	return &BatchAnalyzer{
		analyzer: NewAnalyzer(cfg),
		cfg: cfg,
	}
}

// AnalyzeBatch performs concurrent analysis on multiple articles
func (ba *BatchAnalyzer) AnalyzeBatch(ctx context.Context, req *models.BatchAnalyzeRequest) (*models.BatchAnalyzeResponse, error) {
	if len(req.Articles) == 0 {
		return nil, fmt.Errorf("no articles provided")
	}

	if len(req.Articles) > 100 {
		return nil, fmt.Errorf("batch size exceeds maximum of 100")
	}

	// Create channels for concurrent processing
	resultsChan := make(chan *models.BatchAnalysisResult, len(req.Articles))
	var wg sync.WaitGroup

	// Determine concurrency level (max 10 concurrent goroutines)
	concurrency := len(req.Articles)
	if concurrency > 10 {
		concurrency = 10
	}

	// Create semaphore-like control using a buffered channel
	semaphore := make(chan struct{}, concurrency)

	// Process articles concurrently
	for i, article := range req.Articles {
		wg.Add(1)
		go func(idx int, art models.BatchArticle) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire semaphore
			defer func() { <-semaphore }() // Release semaphore

			// Validate article
			result := &models.BatchAnalysisResult{
				ID: art.ID,
			}

			if len(art.Content) < 10 {
				result.Success = false
				result.Error = "Article must be at least 10 characters"
				resultsChan <- result
				return
			}

			if len(art.Content) > ba.cfg.MaxArticleLength {
				result.Success = false
				result.Error = fmt.Sprintf("Article exceeds maximum length of %d", ba.cfg.MaxArticleLength)
				resultsChan <- result
				return
			}

			// Set default language
			language := art.Language
			if language == "" {
				language = "en"
			}

			// Analyze article
			req := models.AnalyzeRequest{
				Article:  art.Content,
				Language: language,
			}

			resp, err := ba.analyzer.Analyze(ctx, req)
			if err != nil {
				result.Success = false
				result.Error = fmt.Sprintf("Analysis failed: %v", err)
			} else {
				result.Success = true
				result.Data = resp
			}

			resultsChan <- result
		}(i, article)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(resultsChan)

	// Collect results
	var results []models.BatchAnalysisResult
	processedCount := 0
	failedCount := 0

	for result := range resultsChan {
		results = append(results, *result)
		if result.Success {
			processedCount++
		} else {
			failedCount++
		}
	}

	// Return batch response
	return &models.BatchAnalyzeResponse{
		Results:        results,
		ProcessedCount: processedCount,
		FailedCount:    failedCount,
	}, nil
}