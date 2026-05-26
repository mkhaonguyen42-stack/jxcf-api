# API Documentation

## Base URL
```
http://localhost:8080
```

## Authentication
Currently, the API does not require authentication. For production use, consider adding API key authentication.

## Endpoints

### 1. Health Check

#### Request
```
GET /api/health
```

#### Response
```json
{
  "status": "ok"
}
```

### 2. Article Analysis

Analyze an article for summary, keywords, and SEO optimization.

#### Request
```
POST /api/analyze
Content-Type: application/json
```

**Request Body:**
```json
{
  "article": "Your article text here...",
  "language": "zh"
}
```

**Parameters:**
- `article` (string, required): The article content to analyze. Must be between 10 and 100,000 characters.
- `language` (string, optional): Language of the article. Supported values: `"zh"` (Chinese), `"en"` (English). Default: `"zh"`

#### Response
```json
{
  "summary": "AI is transforming business by enhancing operations and decision-making through machine learning.",
  "keywords": [
    {
      "keyword": "artificial intelligence",
      "relevance": 0.95
    },
    {
      "keyword": "machine learning",
      "relevance": 0.89
    },
    {
      "keyword": "business transformation",
      "relevance": 0.82
    },
    {
      "keyword": "data analysis",
      "relevance": 0.76
    },
    {
      "keyword": "automation",
      "relevance": 0.71
    }
  ],
  "meta_description": "AI is transforming business by enhancing operations and decision-making through machine learning.",
  "seo_score": 8.5
}
```

**Response Fields:**
- `summary` (string): SEO-optimized summary of the article (50-160 characters)
- `keywords` (array): Top keywords extracted from the article with relevance scores
  - `keyword` (string): The keyword text
  - `relevance` (number): Relevance score between 0 and 1
- `meta_description` (string): Formatted meta description suitable for HTML meta tags
- `seo_score` (number): Overall SEO optimization score (0-10)

#### Error Response
```json
{
  "error": "Article is required and must be at least 10 characters",
  "code": 400
}
```

## Error Handling

### HTTP Status Codes
- `200 OK`: Request successful
- `400 Bad Request`: Invalid request parameters
- `500 Internal Server Error`: Server-side error

### Error Response Format
```json
{
  "error": "Error description",
  "code": 400
}
```

## Rate Limiting

Currently, no rate limiting is implemented. Consider adding it for production deployments.

## Examples

### Using cURL

```bash
# Health check
curl http://localhost:8080/api/health

# Analyze article in Chinese
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "article": "人工智能正在改变现代业务。全球公司正在采用人工智能来增强其运营。",
    "language": "zh"
  }'

# Analyze article in English
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "article": "Artificial intelligence is transforming modern business. Companies worldwide are adopting AI to enhance their operations.",
    "language": "en"
  }'
```

### Using Python

```python
import requests
import json

url = "http://localhost:8080/api/analyze"
data = {
    "article": "Your article text here...",
    "language": "en"
}

response = requests.post(url, json=data)
result = response.json()

print(f"Summary: {result['summary']}")
print(f"SEO Score: {result['seo_score']}")
for kw in result['keywords']:
    print(f"  - {kw['keyword']}: {kw['relevance']:.2f}")
```

### Using JavaScript/Node.js

```javascript
const axios = require('axios');

const analyzeArticle = async (article, language = 'en') => {
  try {
    const response = await axios.post('http://localhost:8080/api/analyze', {
      article,
      language
    });
    return response.data;
  } catch (error) {
    console.error('Error:', error.response.data);
  }
};

// Usage
analyzeArticle('Your article text...', 'en').then(result => {
  console.log('Summary:', result.summary);
  console.log('Keywords:', result.keywords);
});
```

## Best Practices

1. **Article Length**: Keep articles between 100-5000 characters for optimal performance
2. **Language Setting**: Specify the correct language for better accuracy
3. **Error Handling**: Implement retry logic for production applications
4. **Caching**: Cache results for identical articles to reduce API calls
5. **Monitoring**: Track API response times and error rates

## SEO Score Calculation

The SEO score is calculated based on:
- Summary length (50-160 characters): 30 points
- Keyword count (3-5 keywords): 30 points
- Keyword relevance (average > 0.8): 40 points

Score is normalized to 0-10 scale.
