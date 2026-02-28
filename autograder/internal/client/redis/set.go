package redis

import (
	"context"
	"time"
)

// Set sets a key-value pair
// in Redis with an optional TTL (Time-to-Live).
//
// If ttl is zero, the key will not expire
// and will persist until explicitly deleted.
//
// Returns an error if the operation fails.
func (c *Client) Set(
	ctx context.Context,
	key string,
	value any,
	ttl time.Duration,
) error {
	return c.cli.Set(ctx, key, value, ttl).Err()
}
