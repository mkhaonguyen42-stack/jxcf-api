# Configuration Guide

## Environment Variables

### Required
- `DEEPSEEK_API_KEY` - Your DeepSeek API key

### Optional
- `PORT` - Server port (default: 8080)
- `GIN_MODE` - debug or release (default: debug)
- `DEEPSEEK_BASE_URL` - API endpoint (default: https://api.deepseek.com/v1)
- `DEEPSEEK_MODEL` - Model name (default: deepseek-v4)
- `MAX_ARTICLE_LENGTH` - Max article chars (default: 10000)
- `SUMMARY_MIN_LENGTH` - Min summary chars (default: 50)
- `SUMMARY_MAX_LENGTH` - Max summary chars (default: 160)
- `KEYWORD_COUNT` - Keywords to extract (default: 5)

## Setup

```bash
cp .env.example .env
# Edit .env with your API key
PORT=8080 go run cmd/server/main.go
```
