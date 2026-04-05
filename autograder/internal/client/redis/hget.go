package redis

import "context"

// HGet retrieves the value associated with a field in a hash stored at key.
//
// Returns the value associated with the field as a byte slice,
// or an error if the operation fails
// or if the key or field does not exist.
func (c *Client) HGet(
	ctx context.Context,
	key string,
	field string,
) (string, error) {
	cmd := c.cli.HGet(ctx, key, field)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	return cmd.String(), nil
}
