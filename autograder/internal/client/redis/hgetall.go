package redis

import "context"

// HGetAll retrieves all fields and values of a hash stored at key.
//
// It returns a map of field-value pairs and an error if the operation fails.
func (c *Client) HGetAll(
	ctx context.Context,
	key string,
) (map[string]string, error) {
	return c.cli.HGetAll(ctx, key).Result()
}
