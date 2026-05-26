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

// BatchAnalyzeRequest represents a batch analysis request with multiple articles
type BatchAnalyzeRequest struct {
	Articles []BatchArticle `json:"articles" binding:"required,min=1,max=100"`
}

// BatchArticle represents a single article in a batch request
type BatchArticle struct {
	ID string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required,min=10,max=100000"`
	Language string `json:"language" binding:"omitempty,oneof=zh en"`
}

// BatchAnalyzeResponse represents the batch analysis response
type BatchAnalyzeResponse struct {
	Results []BatchAnalysisResult `json:"results"`
	ProcessedCount int `json:"processed_count"`
	FailedCount int `json:"failed_count"`
}

// BatchAnalysisResult represents the analysis result for a single article
type BatchAnalysisResult struct {
	ID string `json:"id"`
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Data *AnalyzeResponse `json:"data,omitempty"`
}