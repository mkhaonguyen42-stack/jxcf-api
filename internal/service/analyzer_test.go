package service

import (
	"testing"

	"github.com/jxcf/jxcf-api/internal/config"
)

func TestNewAnalyzer(t *testing.T) {
	cfg := &config.Config{
		DeepseekAPIKey: "test-key",
		DeepseekBaseURL: "https://api.deepseek.com/v1",
		DeepseekModel: "deepseek-v4",
		MaxArticleLength: 10000,
		SummaryMinLength: 50,
		SummaryMaxLength: 160,
		KeywordCount: 5,
	}

	analyzer := NewAnalyzer(cfg)
	if analyzer == nil {
		t.Fatal("Expected analyzer to be created")
	}
}
