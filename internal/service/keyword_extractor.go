package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/llm"
	"github.com/jxcf/jxcf-api/internal/models"
)

type KeywordExtractor struct {
	client *llm.DeepSeekClient
	cfg *config.Config
}

func NewKeywordExtractor(client *llm.DeepSeekClient, cfg *config.Config) *KeywordExtractor {
	return &KeywordExtractor{
		client: client,
		cfg: cfg,
	}
}

func (k *KeywordExtractor) Extract(ctx context.Context, article string, language string) ([]models.KeywordItem, error) {
	// Prepare prompt
	var prompt string
	if language == "en" {
		prompt = fmt.Sprintf(`Extract the top %d keywords from the following article. For each keyword, provide a relevance score between 0 and 1.
Return as JSON array with format: [{"keyword": "word", "relevance": 0.95}]

Article:
%s

JSON Response:`, k.cfg.KeywordCount, article)
	} else {
		prompt = fmt.Sprintf(`从以下文章中提取前%d个关键词。对于每个关键词，提供0到1之间的相关性评分。
返回JSON数组格式: [{"keyword": "词", "relevance": 0.95}]

文章:
%s

JSON响应:`, k.cfg.KeywordCount, article)
	}

	messages := []llm.ChatMessage{
		{
			Role: "user",
			Content: prompt,
		},
	}

	response, err := k.client.ChatCompletion(ctx, messages, 300)
	if err != nil {
		return nil, fmt.Errorf("failed to extract keywords: %w", err)
	}

	var keywords []models.KeywordItem
	if err := json.Unmarshal([]byte(response), &keywords); err != nil {
		// Try to parse more flexibly
		keywords, _ = parseKeywordsFallback(response)
	}

	// Sort by relevance
	for i := 0; i < len(keywords)-1; i++ {
		for j := i + 1; j < len(keywords); j++ {
			if keywords[j].Relevance > keywords[i].Relevance {
			keywords[i], keywords[j] = keywords[j], keywords[i]
			}
		}
	}

	return keywords, nil
}

func parseKeywordsFallback(response string) ([]models.KeywordItem, error) {
	var keywords []models.KeywordItem
	lines := strings.Split(response, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var item models.KeywordItem
		if err := json.Unmarshal([]byte(line), &item); err == nil {
			keywords = append(keywords, item)
		}
	}
	return keywords, nil
}
