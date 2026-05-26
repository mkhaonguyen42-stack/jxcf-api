package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/llm"
)

type Summarizer struct {
	client *llm.DeepSeekClient
	cfg *config.Config
}

func NewSummarizer(client *llm.DeepSeekClient, cfg *config.Config) *Summarizer {
	return &Summarizer{
		client: client,
		cfg: cfg,
	}
}

func (s *Summarizer) Generate(ctx context.Context, article string, language string) (string, error) {
	// Prepare prompt
	var prompt string
	if language == "en" {
		prompt = fmt.Sprintf(`Please summarize the following article in 50-160 characters. The summary should be SEO-optimized, clear, and engaging.

Article:
%s

Provide only the summary, without any additional text.`, article)
	} else {
		prompt = fmt.Sprintf(`请总结以下文章，字数在50-160字之间。总结应该是SEO优化的，清晰且引人入胜。

文章:
%s

只提供总结，不需要其他文本。`, article)
	}

	messages := []llm.ChatMessage{
		{
			Role: "user",
			Content: prompt,
		},
	}

	summary, err := s.client.ChatCompletion(ctx, messages, 200)
	if err != nil {
		return "", fmt.Errorf("failed to generate summary: %w", err)
	}

	// Validate and trim summary
	summary = strings.TrimSpace(summary)
	if len(summary) < s.cfg.SummaryMinLength || len(summary) > s.cfg.SummaryMaxLength {
		// Adjust if needed
		if len(summary) > s.cfg.SummaryMaxLength {
			summary = summary[:s.cfg.SummaryMaxLength]
		}
	}

	return summary, nil
}
