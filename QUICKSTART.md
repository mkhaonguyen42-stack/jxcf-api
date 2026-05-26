# Quick Start Guide

## 5 Minute Setup

### 1. Prerequisites
- Go 1.26 or higher
- Git
- DeepSeek API Key

### 2. Clone Repository
```bash
git clone https://github.com/mkhaonguyen42-stack/jxcf-api.git
cd jxcf-api
```

### 3. Configure API Key
```bash
cp .env.example .env
```
Edit `.env` and add your DeepSeek API key

### 4. Run Server
```bash
make run
# or
go run cmd/server/main.go
```

### 5. Test API
```bash
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{"article":"AI is transforming business...","language":"en"}'
```

## Docker Quick Start
```bash
docker build -t jxcf-api .
docker run -p 8080:8080 -e DEEPSEEK_API_KEY=your_key jxcf-api
```

## Next Steps
- Read [API.md](API.md) for detailed endpoints
- Check [DEPLOYMENT.md](DEPLOYMENT.md) for production
- Review [CONFIG.md](CONFIG.md) for configuration
