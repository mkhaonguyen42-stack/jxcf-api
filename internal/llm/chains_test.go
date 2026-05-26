package llm

import (
	"testing"
)

func TestNewPromptTemplate(t *testing.T) {
	template := "Please summarize: {input}"
	variables := []string{"input"}

	pt := NewPromptTemplate(template, variables)
	if pt == nil {
		t.Fatal("Expected PromptTemplate to be created")
	}
}

func TestPromptTemplateFormat(t *testing.T) {
	template := "Summarize this text:"
	variables := []string{"text"}

	pt := NewPromptTemplate(template, variables)
	values := map[string]string{
		"text": "Some important content",
	}

	result := pt.Format(values)
	if result == "" {
		t.Error("Expected formatted result, got empty string")
	}
}

func TestNewLLMChain(t *testing.T) {
	client := NewDeepSeekClient("test-key", "https://api.deepseek.com/v1", "deepseek-v4")
	template := NewPromptTemplate("Summarize:", []string{})

	chain := NewLLMChain(client, template)
	if chain == nil {
		t.Fatal("Expected LLMChain to be created")
	}
}

func TestNewSequentialChain(t *testing.T) {
	client := NewDeepSeekClient("test-key", "https://api.deepseek.com/v1", "deepseek-v4")
	template := NewPromptTemplate("Test:", []string{})

	chain1 := NewLLMChain(client, template)
	chain2 := NewLLMChain(client, template)

	seqChain := NewSequentialChain(chain1, chain2)
	if seqChain == nil {
		t.Fatal("Expected SequentialChain to be created")
	}
}
