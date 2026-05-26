package llm

import (
	"context"
	"fmt"
)

// Chain represents a processing chain in LangChain-like fashion
type Chain interface {
	Run(ctx context.Context, input string) (string, error)
}

// PromptTemplate represents a template for prompts
type PromptTemplate struct {
	template string
	variables []string
}

// NewPromptTemplate creates a new prompt template
func NewPromptTemplate(template string, variables []string) *PromptTemplate {
	return &PromptTemplate{
		template: template,
		variables: variables,
	}
}

// Format formats the template with provided values
func (pt *PromptTemplate) Format(values map[string]string) string {
	result := pt.template
	for _, variable := range pt.variables {
		if value, ok := values[variable]; ok {
			result = fmt.Sprintf("%s%s", result, value)
		}
	}
	return result
}

// LLMChain represents a chain with an LLM client
type LLMChain struct {
	llm *DeepSeekClient
	prompt *PromptTemplate
}

// NewLLMChain creates a new LLM chain
func NewLLMChain(llm *DeepSeekClient, prompt *PromptTemplate) *LLMChain {
	return &LLMChain{
		llm: llm,
		prompt: prompt,
	}
}

// Run executes the chain
func (lc *LLMChain) Run(ctx context.Context, input string) (string, error) {
	fromattedPrompt := lc.prompt.Format(map[string]string{"input": input})
	messages := []ChatMessage{
		{
			Role: "user",
			Content: fromattedPrompt,
		},
	}
	return lc.llm.ChatCompletion(ctx, messages, 500)
}

// SequentialChain represents a sequence of chains
type SequentialChain struct {
	chains []Chain
}

// NewSequentialChain creates a new sequential chain
func NewSequentialChain(chains ...Chain) *SequentialChain {
	return &SequentialChain{
		chains: chains,
	}
}

// Run executes all chains in sequence
func (sc *SequentialChain) Run(ctx context.Context, input string) (string, error) {
	result := input
	for _, chain := range sc.chains {
		var err error
		result, err = chain.Run(ctx, result)
		if err != nil {
			return "", fmt.Errorf("chain execution failed: %w", err)
		}
	}
	return result, nil
}
