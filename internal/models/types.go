package models

// AnalyzeRequest represents the request body for article analysis
type AnalyzeRequest struct {
	Article string `json:"article" binding:"required,min=10,max=100000"`
	Language string `json:"language" binding:"omitempty,oneof=zh en"`
}

// AnalyzeResponse represents the response with summary and keywords
type AnalyzeResponse struct {
	Summary string `json:"summary"`
	Keywords []KeywordItem `json:"keywords"`
	MetaDescription string `json:"meta_description"`
	SEOScore float64 `json:"seo_score"`
}

// KeywordItem represents a single keyword with relevance score
type KeywordItem struct {
	Keyword string `json:"keyword"`
	Relevance float64 `json:"relevance"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
	Code int `json:"code"`
}
