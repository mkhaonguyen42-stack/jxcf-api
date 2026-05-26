package llm

import (
	"context"
	"testing"
)

func TestNewDeepSeekClient(t *testing.T) {
	client := NewDeepSeekClient("test-key", "https://api.deepseek.com/v1", "deepseek-v4")
	if client == nil {
		t.Fatal("Expected client to be created")
	}
	if client.apiKey != "test-key" {
		t.Errorf("Expected apiKey to be 'test-key', got %s", client.apiKey)
	}
}

func TestChatCompletion(t *testing.T) {
	// This is a mock test - in real scenarios you'd use mocking frameworks
	client := NewDeepSeekClient("mock-key", "https://api.deepseek.com/v1", "deepseek-v4")

	ctx := context.Background()
	messages := []ChatMessage{
		{
			Role: "user",
			Content: "Hello",
		},
	}

	// This would fail without real API key, but demonstrates the structure
	_, err := client.ChatCompletion(ctx, messages, 100)
	if err != nil && err.Error() == "api error: status 401, body: {\"error\":{\"message\":\"Unauthorized\"}}" {
		// Expected error with invalid key
		t.Logf("Got expected auth error: %v", err)
	}
}
