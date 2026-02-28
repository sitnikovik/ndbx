package redis

import (
	"context"
	"time"
)

// TTL returns the time to live of the given key.
//
// It returns the remaining time to live of the key in seconds.
// If the key does not exist, it returns a negative duration and no error.
// If the key exists but has no associated expiration time, it returns a negative duration and no error.
// If there is an error while fetching the TTL, it returns an error.
func (c *Client) TTL(
	ctx context.Context,
	key string,
) (time.Duration, error) {
	return c.cli.TTL(ctx, key).Result()
}
