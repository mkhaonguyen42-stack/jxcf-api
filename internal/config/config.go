package config

import (
	"os"
	"strconv"
)

type Config struct {
	// API Configuration
	Port int
	GinMode string

	// DeepSeek Configuration
	DeepseekAPIKey string
	DeepseekBaseURL string
	DeepseekModel string

	// Text Processing
	MaxArticleLength int
	SummaryMinLength int
	SummaryMaxLength int
	KeywordCount int
}

func Load() *Config {
	return &Config{
		Port: getEnvInt("PORT", 8080),
		GinMode: getEnv("GIN_MODE", "debug"),

		DeepseekAPIKey: getEnv("DEEPSEEK_API_KEY", ""),
		DeepseekBaseURL: getEnv("DEEPSEEK_BASE_URL", "https://api.deepseek.com/v1"),
		DeepseekModel: getEnv("DEEPSEEK_MODEL", "deepseek-v4"),

		MaxArticleLength: getEnvInt("MAX_ARTICLE_LENGTH", 10000),
		SummaryMinLength: getEnvInt("SUMMARY_MIN_LENGTH", 50),
		SummaryMaxLength: getEnvInt("SUMMARY_MAX_LENGTH", 160),
		KeywordCount: getEnvInt("KEYWORD_COUNT", 5),
	}
}

func getEnv(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultVal
}
