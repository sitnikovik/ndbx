package redis

import "context"

// Ping checks the connection to the Redis server.
//
// It returns an error if the connection is not healthy.
// This is a simple way to verify that the Redis server is reachable and responsive.
func (c *Client) Ping(ctx context.Context) error {
	return c.cli.Ping(ctx).Err()
}
