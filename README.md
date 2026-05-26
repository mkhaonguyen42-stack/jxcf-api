# jxcf-api

A Go REST API for article summarization and keyword extraction powered by DeepSeek v4 and LangChain.

## Features

- **Article Summarization**: Generate SEO-optimized summaries (50-160 characters) using DeepSeek v4
- **Keyword Extraction**: Extract top 3-5 keywords with relevance scores
- **SEO Optimization**: Automatic meta description formatting and SEO scoring
- **Multi-language Support**: Support for Chinese (zh) and English (en)
- **RESTful API**: Simple HTTP API with JSON request/response

## Requirements

- Go 1.26+
- DeepSeek API Key (v4 model)

## Installation

```bash
git clone https://github.com/mkhaonguyen42-stack/jxcf-api.git
cd jxcf-api
go mod download
```

## Configuration

Create a `.env` file from `.env.example`:

```bash
cp .env.example .env
```

Edit `.env` and set your DeepSeek API key:

```env
DEEPSEEK_API_KEY=your_api_key_here
PORT=8080
```

## Running

```bash
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Health Check

```
GET /api/health
```

Response:
```json
{
  "status": "ok"
}
```

### Article Analysis

```
POST /api/analyze
Content-Type: application/json

{
  "article": "Your article text here...",
  "language": "zh"
}
```

Response:
```json
{
  "summary": "SEO-optimized summary",
  "keywords": [
    {
      "keyword": "keyword1",
      "relevance": 0.95
    },
    {
      "keyword": "keyword2",
      "relevance": 0.88
    }
  ],
  "meta_description": "SEO meta description",
  "seo_score": 8.5
}
```

## Architecture

```
jxcf-api/
├── cmd/server/          # Application entrypoint
├── internal/
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP handlers
│   ├── llm/             # DeepSeek API client
│   ├── models/          # Data structures
│   └── service/         # Business logic
└── pkg/
    └── seo/             # SEO utilities
```

## Development

Run tests:
```bash
go test ./...
```

Format code:
```bash
go fmt ./...
```

Lint:
```bash
golangci-lint run
```

## License

MIT License
