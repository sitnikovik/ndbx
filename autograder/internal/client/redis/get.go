package redis

import "context"

// Get retrieves the value associated with a key from Redis.
//
// Returns the value associated with the key as a byte slice,
// or an error if the operation fails
// or if the key does not exist.
func (c *Client) Get(
	ctx context.Context,
	key string,
) ([]byte, error) {
	return c.cli.Get(ctx, key).Bytes()
}
