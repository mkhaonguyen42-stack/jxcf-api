package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenBucket struct {
	capacity    float64
	tokens      float64
	refillRate  float64
	lastRefill  time.Time
	mu          sync.Mutex
}

func NewTokenBucket(capacity float64, tokensPerSecond float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: tokensPerSecond,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Allow(tokens float64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens = min(tb.capacity, tb.tokens+elapsed*tb.refillRate)
	tb.lastRefill = now

	if tb.tokens >= tokens {
		tb.tokens -= tokens
		return true
	}
	return false
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

type RateLimiter struct {
	buckets    map[string]*TokenBucket
	capacity   float64
	refillRate float64
	mu         sync.Mutex
}

func NewRateLimiter(requestsPerMinute int, burstSize int) *RateLimiter {
	tokensPerSecond := float64(requestsPerMinute) / 60.0
	capacity := float64(burstSize)
	if capacity < tokensPerSecond {
		capacity = tokensPerSecond * 2
	}
	return &RateLimiter{
		buckets:    make(map[string]*TokenBucket),
		capacity:   capacity,
		refillRate: tokensPerSecond,
	}
}

func (rl *RateLimiter) Allow(clientIP string) bool {
	rl.mu.Lock()
	bucket, exists := rl.buckets[clientIP]
	if !exists {
		bucket = NewTokenBucket(rl.capacity, rl.refillRate)
		rl.buckets[clientIP] = bucket
	}
	rl.mu.Unlock()

	return bucket.Allow(1)
}

func getClientIP(c *gin.Context) string {
	ip := c.GetHeader("X-Forwarded-For")
	if ip != "" {
		if comma := net.SplitHostPort(ip); comma != "" {
			ip = comma
		}
		return ip
	}

	ip = c.GetHeader("X-Real-IP")
	if ip != "" {
		return ip
	}

	return c.ClientIP()
}

func RateLimitMiddleware(limiter *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := getClientIP(c)

		if !limiter.Allow(clientIP) {
			c.Header("Retry-After", "60")
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
				"code":  http.StatusTooManyRequests,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
