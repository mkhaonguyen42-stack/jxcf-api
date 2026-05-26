package service

import (
	"context"
	"testing"

	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/models"
)

func TestBatchAnalyzerCreate(t *testing.T) {
	cfg := &config.Config{
		MaxArticleLength: 100000,
	}

	ba := NewBatchAnalyzer(cfg)
	if ba == nil {
		t.Error("Expected BatchAnalyzer to be created")
	}

	if ba.cfg != cfg {
		t.Error("Config not set correctly")
	}
}

func TestBatchAnalyzerWithEmptyArticles(t *testing.T) {
	cfg := &config.Config{
		MaxArticleLength: 100000,
	}

	ba := NewBatchAnalyzer(cfg)
	req := &models.BatchAnalyzeRequest{
		Articles: []models.BatchArticle{},
	}

	_, err := ba.AnalyzeBatch(context.Background(), req)
	if err == nil {
		t.Error("Expected error for empty articles")
	}
}

func TestBatchAnalyzerWithTooManyArticles(t *testing.T) {
	cfg := &config.Config{
		MaxArticleLength: 100000,
	}

	ba := NewBatchAnalyzer(cfg)
	articles := make([]models.BatchArticle, 101)
	for i := 0; i < 101; i++ {
		articles[i] = models.BatchArticle{
			ID:       "test-id",
			Content:  "Test article content for testing",
			Language: "en",
		}
	}

	req := &models.BatchAnalyzeRequest{
		Articles: articles,
	}

	_, err := ba.AnalyzeBatch(context.Background(), req)
	if err == nil {
		t.Error("Expected error for too many articles")
	}
}

func TestBatchAnalyzerWithValidArticles(t *testing.T) {
	cfg := &config.Config{
		MaxArticleLength: 100000,
		DeepseekAPIKey:   "test-key",
		DeepseekBaseURL:  "https://api.deepseek.com",
		DeepseekModel:    "deepseek-chat",
	}

	ba := NewBatchAnalyzer(cfg)

	// Create test articles with short content for quick validation
	articles := []models.BatchArticle{
		{
			ID:       "article-1",
			Content:  "This is a test article with minimum required content",
			Language: "en",
		},
		{
			ID:       "article-2",
			Content:  "Another test article content here",
			Language: "en",
		},
	}

	req := &models.BatchAnalyzeRequest{
		Articles: articles,
	}

	// This will make actual API calls if credentials are valid
	// For testing purposes, we're checking that the batch response structure is created
	resp, err := ba.AnalyzeBatch(context.Background(), req)
	if resp == nil {
		t.Error("Expected batch response to be created")
	}

	if len(resp.Results) != len(articles) {
		t.Errorf("Expected %d results, got %d", len(articles), len(resp.Results))
	}
}

func TestBatchAnalyzerWithTooShortArticle(t *testing.T) {
	cfg := &config.Config{
		MaxArticleLength: 100000,
	}

	ba := NewBatchAnalyzer(cfg)

	articles := []models.BatchArticle{
		{
			ID:       "article-1",
			Content:  "short",
			Language: "en",
		},
	}

	req := &models.BatchAnalyzeRequest{
		Articles: articles,
	}

	resp, err := ba.AnalyzeBatch(context.Background(), req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resp == nil {
		t.Error("Expected batch response")
		return
	}

	if resp.FailedCount != 1 {
		t.Errorf("Expected 1 failed article, got %d", resp.FailedCount)
	}

	if resp.Results[0].Success {
		t.Error("Expected article to fail validation")
	}
}

func TestBatchAnalyzerWithExceedingMaxLength(t *testing.T) {
	cfg := &config.Config{
		MaxArticleLength: 100,
	}

	ba := NewBatchAnalyzer(cfg)

	content := ""
	for i := 0; i < 50; i++ {
		content += "This is a long article content that exceeds the maximum allowed length. "
	}

	articles := []models.BatchArticle{
		{
			ID:       "article-1",
			Content:  content,
			Language: "en",
		},
	}

	req := &models.BatchAnalyzeRequest{
		Articles: articles,
	}

	resp, err := ba.AnalyzeBatch(context.Background(), req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resp == nil {
		t.Error("Expected batch response")
		return
	}

	if resp.FailedCount != 1 {
		t.Errorf("Expected 1 failed article, got %d", resp.FailedCount)
	}

	if resp.Results[0].Success {
		t.Error("Expected article to fail length validation")
	}
}

func TestBatchAnalyzerConcurrency(t *testing.T) {
	cfg := &config.Config{
		MaxArticleLength: 100000,
		DeepseekAPIKey:   "test-key",
		DeepseekBaseURL:  "https://api.deepseek.com",
		DeepseekModel:    "deepseek-chat",
	}

	ba := NewBatchAnalyzer(cfg)

	// Create 20 test articles to verify concurrent processing
	articles := make([]models.BatchArticle, 20)
	for i := 0; i < 20; i++ {
		articles[i] = models.BatchArticle{
			ID:       "article-" + string(rune(i)),
			Content:  "This is test article number for concurrent processing",
			Language: "en",
		}
	}

	req := &models.BatchAnalyzeRequest{
		Articles: articles,
	}

	resp, err := ba.AnalyzeBatch(context.Background(), req)
	if resp == nil {
		t.Error("Expected batch response")
	}

	if len(resp.Results) != 20 {
		t.Errorf("Expected 20 results, got %d", len(resp.Results))
	}
}
