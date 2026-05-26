package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/llm"
)

type ChainSummarizer struct {
	client *llm.DeepSeekClient
	cfg *config.Config
	chain llm.Chain
}

// NewChainSummarizer creates a summarizer using LLM chain
func NewChainSummarizer(client *llm.DeepSeekClient, cfg *config.Config) *ChainSummarizer {
	// Create prompt template for summarization
	promptTemplate := llm.NewPromptTemplate(
		"Summarize the following article in 50-160 characters:\n",
		[]string{"input"},
	)

	chain := llm.NewLLMChain(client, promptTemplate)

	return &ChainSummarizer{
		client: client,
		cfg: cfg,
		chain: chain,
	}
}

// GenerateWithChain generates summary using the LangChain approach
func (cs *ChainSummarizer) GenerateWithChain(ctx context.Context, article string, language string) (string, error) {
	var prompt string
	if language == "en" {
		prompt = fmt.Sprintf("Summarize in 50-160 characters:\n%s", article)
	} else {
		prompt = fmt.Sprintf("总结以下文章，字数50-160字：\n%s", article)
	}

	summary, err := cs.chain.Run(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("chain execution failed: %w", err)
	}

	return strings.TrimSpace(summary), nil
}
