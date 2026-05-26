package seo

import (
	"testing"

	"github.com/jxcf/jxcf-api/internal/models"
)

func TestFormatMetaDescription(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Short description",
			input:    "Short summary",
			expected: "Short summary",
		},
		{
			name:     "Long description",
			input:    "This is a very long description that exceeds the maximum character limit and should be truncated with ellipsis added to indicate continuation of the text",
			expected: "This is a very long description that exceeds the maximum character limit and should be truncated with ellipsis added to indicate continuation...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatMetaDescription(tt.input)
			if result != tt.expected {
				t.Errorf("Got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestCalculateSEOScore(t *testing.T) {
	keywords := []models.KeywordItem{
		{Keyword: "keyword1", Relevance: 0.95},
		{Keyword: "keyword2", Relevance: 0.87},
		{Keyword: "keyword3", Relevance: 0.76},
	}

	summary := "This is a SEO optimized summary that is between 50 and 160 characters long."
	score := CalculateSEOScore(summary, keywords)

	if score < 0 || score > 10 {
		t.Errorf("SEO score out of range: %f", score)
	}
	t.Logf("Calculated SEO score: %f", score)
}

func TestValidateSEOContent(t *testing.T) {
	tests := []struct {
		name       string
		summary    string
		keywords   []models.KeywordItem
		expectErrs bool
	}{
		{
		name:    "Valid content",
		summary: "This is a SEO optimized summary that is between 50 and 160 characters long.",
		keywords: []models.KeywordItem{
			{Keyword: "keyword1", Relevance: 0.95},
			{Keyword: "keyword2", Relevance: 0.87},
			{Keyword: "keyword3", Relevance: 0.76},
		},
		expectErrs: false,
		},
		{
		name:    "Summary too short",
		summary: "Too short",
		keywords: []models.KeywordItem{
			{Keyword: "keyword1", Relevance: 0.95},
		},
		expectErrs: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := ValidateSEOContent(tt.summary, tt.keywords)
			if (len(issues) > 0) != tt.expectErrs {
				t.Errorf("Expected errors: %v, got: %v", tt.expectErrs, issues)
			}
		})
	}
}
