package service

import (
	"context"
	"fmt"

	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/llm"
	"github.com/jxcf/jxcf-api/internal/models"
	"github.com/jxcf/jxcf-api/pkg/seo"
)

type Analyzer struct {
	summarizer *Summarizer
	keywordExtractor *KeywordExtractor
	cfg *config.Config
}

func NewAnalyzer(cfg *config.Config) *Analyzer {
	client := llm.NewDeepSeekClient(
		cfg.DeepseekAPIKey,
		cfg.DeepseekBaseURL,
		cfg.DeepseekModel,
	)

	return &Analyzer{
		summarizer: NewSummarizer(client, cfg),
		keywordExtractor: NewKeywordExtractor(client, cfg),
		cfg: cfg,
	}
}

func (a *Analyzer) Analyze(ctx context.Context, req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {
	// Validate article length
	if len(req.Article) > a.cfg.MaxArticleLength {
		return nil, fmt.Errorf("article exceeds maximum length of %d characters", a.cfg.MaxArticleLength)
	}

	// Generate summary
	summary, err := a.summarizer.Generate(ctx, req.Article, req.Language)
	if err != nil {
		return nil, fmt.Errorf("failed to generate summary: %w", err)
	}

	// Extract keywords
	keywords, err := a.keywordExtractor.Extract(ctx, req.Article, req.Language)
	if err != nil {
		return nil, fmt.Errorf("failed to extract keywords: %w", err)
	}

	// Calculate SEO score
	seoScore := seo.CalculateSEOScore(summary, keywords)

	// Format response
	resp := &models.AnalyzeResponse{
		Summary: summary,
		Keywords: keywords,
		MetaDescription: seo.FormatMetaDescription(summary),
		SEOScore: seoScore,
	}

	return resp, nil
}
