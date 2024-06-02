package ratelimit

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter_Allow(t *testing.T) {
	// Start a miniredis server for testing
	srv := miniredis.RunT(t)

	// Create a redis client connected to the miniredis server
	rdb := redis.NewClient(&redis.Options{
		Addr: srv.Addr(),
	})

	// Create a RateLimiter instance
	rl := NewRateLimiter(rdb, 5, time.Minute, "test_prefix")

	// Define a test context
	ctx := context.Background()

	// Test case: Allow requests within the rate limit
	identifier := "user:123"
	for i := 0; i < 5; i++ {
		allowed, err := rl.Allow(ctx, identifier)
		assert.NoError(t, err)
		assert.True(t, allowed)
	}

	// Test case: Rate limit exceeded
	allowed, err := rl.Allow(ctx, identifier)
	assert.NoError(t, err)
	assert.False(t, allowed)

	// Test case: Reset rate limit after expiration
	time.Sleep(time.Minute + time.Second)
	allowed, err = rl.Allow(ctx, identifier)
	assert.NoError(t, err)
	assert.True(t, allowed)
}

func TestRateLimiter_getKey(t *testing.T) {
	rl := NewRateLimiter(nil, 5, time.Minute, "test_prefix")
	key := rl.getKey("user:123")
	assert.Equal(t, "test_prefix:user:123", key)
}
func TestRateLimiter_Allow_WithDifferentPrefixes(t *testing.T) {
	// Start a miniredis server for testing
	srv, err := miniredis.Run()
	if err != nil {
		t.Fatalf("Failed to start miniredis server: %v", err)
	}
	defer srv.Close()

	// Create a redis client connected to the miniredis server
	rdb := redis.NewClient(&redis.Options{
		Addr: srv.Addr(),
	})

	// Define test cases with different prefixes
	testCases := []struct {
		prefix     string
		identifier string
	}{
		{"prefix1", "user:123"},
		{"prefix2", "user:456"},
		{"prefix3", "user:789"},
		{"prefix4", "user:101112"},
		{"prefix5", "user:131415"},
	}

	for _, tc := range testCases {
		// Create a RateLimiter instance with the current test case prefix
		rl := NewRateLimiter(rdb, 5, time.Minute, tc.prefix)

		// Define a test context
		ctx := context.Background()

		// Test case: Allow requests within the rate limit
		for i := 0; i < 5; i++ {
			allowed, err := rl.Allow(ctx, tc.identifier)
			if err != nil {
				t.Fatalf("Error allowing request with prefix '%s' and identifier '%s': %v", tc.prefix, tc.identifier, err)
			}
			if !allowed {
				t.Fatalf("Expected request with prefix '%s' and identifier '%s' to be allowed, but it was rate limited", tc.prefix, tc.identifier)
			}
		}

		// Test case: Rate limit exceeded
		allowed, err := rl.Allow(ctx, tc.identifier)
		if err != nil {
			t.Fatalf("Error allowing request with prefix '%s' and identifier '%s': %v", tc.prefix, tc.identifier, err)
		}
		if allowed {
			t.Fatalf("Expected request with prefix '%s' and identifier '%s' to be rate limited, but it was allowed", tc.prefix, tc.identifier)
		}
	}
}
