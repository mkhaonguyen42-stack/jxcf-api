package middleware

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tb := NewTokenBucket(10, 1)

	if !tb.Allow(1) {
		t.Error("Expected Allow to return true for first request")
	}

	if !tb.Allow(1) {
		t.Error("Expected Allow to return true for second request")
	}

	for i := 0; i < 8; i++ {
		tb.Allow(1)
	}

	if tb.Allow(1) {
		t.Error("Expected Allow to return false when capacity exceeded")
	}

	time.Sleep(1 * time.Second)

	if !tb.Allow(1) {
		t.Error("Expected Allow to return true after refill")
	}
}

func TestRateLimiter(t *testing.T) {
	limiter := NewRateLimiter(60, 20)

	clientIP := "192.168.1.1"

	for i := 0; i < 20; i++ {
		if !limiter.Allow(clientIP) {
			t.Errorf("Expected Allow to return true for request %d", i+1)
		}
	}

	if limiter.Allow(clientIP) {
		t.Error("Expected Allow to return false when burst exceeded")
	}

	anotherIP := "192.168.1.2"
	if !limiter.Allow(anotherIP) {
		t.Error("Expected different IP to have independent limit")
	}
}

func TestRateLimiterRefill(t *testing.T) {
	limiter := NewRateLimiter(1, 1)

	clientIP := "192.168.1.1"

	if !limiter.Allow(clientIP) {
		t.Error("Expected first request to be allowed")
	}

	if limiter.Allow(clientIP) {
		t.Error("Expected second request to be denied")
	}

	time.Sleep(1100 * time.Millisecond)

	if !limiter.Allow(clientIP) {
		t.Error("Expected third request to be allowed after refill")
	}
}
