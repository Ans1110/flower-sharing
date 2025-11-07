package middlewares

import (
	"flower-backend/utils"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type rateLimiter struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
	limit    int
	window   time.Duration
	cleanup  *time.Ticker
}

var (
	limiter *rateLimiter
	once    sync.Once
)

// RateLimiter returns a Gin middleware that limits requests per IP
// 1-minute time window for request limiting
// Allow a maximum of 60 requests per window per IP
// Use the latest standard rate-limit headers (RFC 7239)
// Disable deprecated X-RateLimit headers
func RateLimiter() gin.HandlerFunc {
	once.Do(func() {
		limiter = &rateLimiter{
			requests: make(map[string][]time.Time),
			limit:    60,
			window:   1 * time.Minute,
			cleanup:  time.NewTicker(5 * time.Minute),
		}

		// Start cleanup goroutine to remove old entries
		go limiter.cleanupOldEntries()
	})

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		// Clean up old requests outside the time window
		limiter.cleanupIP(ip, now)

		// Count requests in the current window
		count := len(limiter.requests[ip])

		if count >= limiter.limit {
			// Calculate reset time (when the oldest request expires)
			oldestRequest := limiter.requests[ip][0]
			resetTime := oldestRequest.Add(limiter.window)

			// Set standard rate-limit headers (RFC 7239)
			c.Header("RateLimit-Limit", "60")
			c.Header("RateLimit-Remaining", "0")
			c.Header("RateLimit-Reset", utils.FormatUnixTime(resetTime.Unix()))

			// Retry-After header (seconds until reset)
			retryAfter := int(time.Until(resetTime).Seconds())
			if retryAfter < 0 {
				retryAfter = 0
			}
			c.Header("Retry-After", strconv.Itoa(retryAfter))

			zap.L().Warn("rate limit exceeded",
				zap.String("ip", ip),
				zap.Int("count", count),
				zap.Time("reset_time", resetTime),
			)

			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":       "Too many requests",
				"message":     "Rate limit exceeded. Please try again later.",
				"retry_after": retryAfter,
			})
			return
		}

		// Add current request
		limiter.requests[ip] = append(limiter.requests[ip], now)

		// Calculate remaining requests
		remaining := limiter.limit - count - 1

		// Calculate reset time
		var resetTime time.Time
		if len(limiter.requests[ip]) > 0 {
			oldestRequest := limiter.requests[ip][0]
			resetTime = oldestRequest.Add(limiter.window)
		} else {
			resetTime = now.Add(limiter.window)
		}

		// Set standard rate-limit headers (RFC 7239)
		c.Header("RateLimit-Limit", "60")
		c.Header("RateLimit-Remaining", strconv.Itoa(remaining))
		c.Header("RateLimit-Reset", utils.FormatUnixTime(resetTime.Unix()))

		c.Next()
	}
}

// cleanupIP removes requests older than the time window for a specific IP
func (rl *rateLimiter) cleanupIP(ip string, now time.Time) {
	requests := rl.requests[ip]
	cutoff := now.Add(-rl.window)

	// Find the first request that's still within the window
	validStart := 0
	for i, reqTime := range requests {
		if reqTime.After(cutoff) {
			validStart = i
			break
		}
	}

	// Keep only valid requests
	if validStart > 0 {
		rl.requests[ip] = requests[validStart:]
	} else if len(requests) > 0 && requests[0].Before(cutoff) {
		// All requests are old, clear the slice
		rl.requests[ip] = []time.Time{}
	}
}

// cleanupOldEntries periodically removes IPs with no recent requests
func (rl *rateLimiter) cleanupOldEntries() {
	for range rl.cleanup.C {
		rl.mu.Lock()
		now := time.Now()
		cutoff := now.Add(-rl.window)

		for ip, requests := range rl.requests {
			if len(requests) == 0 {
				delete(rl.requests, ip)
				continue
			}

			// Check if all requests are old
			if len(requests) > 0 && requests[len(requests)-1].Before(cutoff) {
				delete(rl.requests, ip)
			}
		}

		rl.mu.Unlock()
	}
}
