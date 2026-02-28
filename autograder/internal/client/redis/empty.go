package redis

import "context"

// Empty checks if there are no keys in Redis.
//
// Returns true if there are no keys in Redis, false if there are keys,
// and an error if the operation fails.
func (c *Client) Empty(ctx context.Context) (bool, error) {
	keys, err := c.cli.Keys(ctx, "*").Result()
	if err != nil {
		return false, err
	}
	return len(keys) == 0, nil
}
