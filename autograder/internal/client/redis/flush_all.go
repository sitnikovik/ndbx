package redis

import "context"

// FlushAll clears all data from the Redis server.
//
// This is used to ensure a clean state before running each check in the autograder.
// It removes ALL keys from ALL databases, so use with caution.
func (c *Client) FlushAll(ctx context.Context) error {
	return c.cli.FlushAll(ctx).Err()
}
