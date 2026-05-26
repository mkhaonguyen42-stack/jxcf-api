# Deployment Guide

## Local Development

### Prerequisites
- Go 1.26+
- DeepSeek API key

### Setup

1. Clone the repository
```bash
git clone https://github.com/mkhaonguyen42-stack/jxcf-api.git
cd jxcf-api
```

2. Configure environment
```bash
cp .env.example .env
# Edit .env and add your DeepSeek API key
```

3. Run the server
```bash
make run
# or
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`

## Docker Deployment

### Build and Run

1. Build the Docker image
```bash
docker build -t jxcf-api:latest .
```

2. Run with Docker
```bash
docker run -p 8080:8080 \
  -e DEEPSEEK_API_KEY=your_api_key \
  jxcf-api:latest
```

### Docker Compose

1. Set environment variables
```bash
echo "DEEPSEEK_API_KEY=your_api_key" > .env
```

2. Start services
```bash
docker-compose up -d
```

3. Check logs
```bash
docker-compose logs -f jxcf-api
```

4. Stop services
```bash
docker-compose down
```

## Kubernetes Deployment

### Prerequisites
- kubectl configured
- Docker image pushed to registry

### Deploy

1. Create namespace
```bash
kubectl create namespace jxcf
```

2. Create secret for API key
```bash
kubectl create secret generic deepseek-secret \
  --from-literal=api-key=your_api_key \
  -n jxcf
```

3. Deploy using kubectl
```bash
kubectl apply -f deployment.yaml -n jxcf
```

4. Verify deployment
```bash
kubectl get pods -n jxcf
kubectl get svc -n jxcf
```

## API Testing

### Health Check
```bash
curl http://localhost:8080/api/health
```

### Article Analysis
```bash
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "article": "Your article text here...",
    "language": "en"
  }'
```

## Performance Tuning

- Adjust `MAX_ARTICLE_LENGTH` for memory optimization
- Set `GIN_MODE=release` for production
- Configure `DEEPSEEK_MODEL` based on latency requirements

## Monitoring

- Health endpoint: `/api/health`
- Check container logs: `docker logs jxcf-api`
- Monitor API response times with your APM tool

## Troubleshooting

### Connection Errors
- Verify DEEPSEEK_API_KEY is set correctly
- Check network connectivity to DeepSeek API
- Verify DEEPSEEK_BASE_URL is correct

### Timeout Issues
- Increase timeout values for large articles
- Check DeepSeek service status
- Monitor system resources
