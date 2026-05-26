package seo

import (
	"fmt"
	"strings"

	"github.com/jxcf/jxcf-api/internal/models"
)

// FormatMetaDescription formats the summary as a SEO-optimized meta description
func FormatMetaDescription(summary string) string {
	// Ensure it's within Google's recommended range (50-160 characters)
	if len(summary) > 160 {
		// Truncate and add ellipsis
		return summary[:157] + "..."
	}
	return summary
}

// CalculateSEOScore calculates an overall SEO optimization score
func CalculateSEOScore(summary string, keywords []models.KeywordItem) float64 {
	score := 0.0

	// Check summary length (optimal: 50-160 characters)
	summaryLen := len(summary)
	if summaryLen >= 50 && summaryLen <= 160 {
		score += 30.0
	} else if summaryLen >= 40 && summaryLen <= 170 {
		score += 20.0
	} else if summaryLen > 0 {
		score += 10.0
	}

	// Check keyword quality
	if len(keywords) >= 3 && len(keywords) <= 5 {
		score += 30.0
	} else if len(keywords) > 0 {
		score += 20.0
	}

	// Check keyword relevance
	avgRelevance := 0.0
	for _, kw := range keywords {
		avgRelevance += kw.Relevance
	}
	if len(keywords) > 0 {
		avgRelevance /= float64(len(keywords))
	}

	if avgRelevance > 0.8 {
		score += 40.0
	} else if avgRelevance > 0.6 {
		score += 30.0
	} else if avgRelevance > 0.4 {
		score += 20.0
	} else {
		score += 10.0
	}

	return score / 10.0 // Normalize to 0-10 scale
}

// ValidateSEOContent checks if content meets SEO standards
func ValidateSEOContent(summary string, keywords []models.KeywordItem) []string {
	var issues []string

	// Validate summary length
	if len(summary) < 50 {
		issues = append(issues, "Summary is too short (minimum 50 characters)")
	}
	if len(summary) > 160 {
		issues = append(issues, "Summary is too long (maximum 160 characters)")
	}

	// Validate keyword count
	if len(keywords) < 3 {
		issues = append(issues, fmt.Sprintf("Not enough keywords (found %d, need at least 3)", len(keywords)))
	}
	if len(keywords) > 5 {
		issues = append(issues, fmt.Sprintf("Too many keywords (found %d, maximum 5)", len(keywords)))
	}

	// Check for empty keywords
	for _, kw := range keywords {
		if strings.TrimSpace(kw.Keyword) == "" {
			issues = append(issues, "Found empty keyword")
			break
		}
	}

	// Check keyword relevance
	for _, kw := range keywords {
		if kw.Relevance < 0.4 {
			issues = append(issues, fmt.Sprintf("Keyword '%s' has low relevance score (%.2f)", kw.Keyword, kw.Relevance))
		}
	}

	return issues
}
