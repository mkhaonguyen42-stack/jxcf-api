package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DeepSeekClient struct {
	apiKey string
	baseURL string
	model string
	httpClient *http.Client
}

type ChatMessage struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model string `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Temperature float64 `json:"temperature,omitempty"`
	MaxTokens int `json:"max_tokens,omitempty"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
	} `json:"usage"`
}

func NewDeepSeekClient(apiKey, baseURL, model string) *DeepSeekClient {
	return &DeepSeekClient{
		apiKey: apiKey,
		baseURL: baseURL,
		model: model,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *DeepSeekClient) ChatCompletion(ctx context.Context, messages []ChatMessage, maxTokens int) (string, error) {
	req := ChatRequest{
		Model: c.model,
		Messages: messages,
		Temperature: 0.7,
		MaxTokens: maxTokens,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return chatResp.Choices[0].Message.Content, nil
}
