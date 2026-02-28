package redis

import "context"

// Has checks if a key exists in Redis.
//
// Returns true if the key exists, false if it does not exist,
// and an error if the operation fails.
func (c *Client) Has(ctx context.Context, key string) (bool, error) {
	n, err := c.cli.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}
