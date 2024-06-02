package ratelimit

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RateLimiter is an interface that defines a method for rate limiting.
type RateLimiter interface {
	// Allow checks if the given identifier can perform an action based on the rate limit.
	Allow(ctx context.Context, key string) (bool, error)
}

// RateLimit implements the RateLimiter interface using Redis as the storage backend.
type RateLimit struct {
	client *redis.Client
	rate   int
	per    time.Duration
	prefix string
}

// NewRateLimiter creates a new RateLimit instance with the given Redis client, rate limit,
// time period, and prefix for keys.
func NewRateLimiter(client *redis.Client, rate int, per time.Duration, prefix string) *RateLimit {
	return &RateLimit{
		client: client,
		rate:   rate,
		per:    per,
		prefix: prefix,
	}
}

// getKey generates a unique key for the given identifier using the prefix.
func (rl *RateLimit) getKey(identifier string) string {
	return fmt.Sprintf("%s:%s", rl.prefix, identifier)
}

// Allow checks if the given identifier can perform an action based on the rate limit.
// It uses Redis sorted set to store timestamps of requests and removes expired ones.
// It returns true if the request is allowed, false otherwise, along with any error encountered.
func (rl *RateLimit) Allow(ctx context.Context, identifier string) (bool, error) {
	key := rl.getKey(identifier)
	now := time.Now()

	pipeline := rl.client.TxPipeline()

	pipeline.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", now.Add(-rl.per).UnixNano()))

	pipeline.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now.UnixNano()),
		Member: now.UnixNano(),
	})

	pipeline.ZCard(ctx, key)
	pipeline.Expire(ctx, key, rl.per)

	cmders, err := pipeline.Exec(ctx)

	if err != nil {
		return false, err
	}

	zCardCmd := cmders[2].(*redis.IntCmd)

	count := zCardCmd.Val()

	return count <= int64(rl.rate), nil
}
