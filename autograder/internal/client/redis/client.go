package redis

import (
	"context"

	redis "github.com/redis/go-redis/v9"

	config "github.com/sitnikovik/ndbx/autograder/internal/config/redis"
)

// Client is a wrapper around the Redis client
// that provides methods for interacting with the Redis server.
type Client struct {
	cli *redis.Client
}

// NewClient creates a new Redis client using the provided configuration.
func NewClient(cfg config.Config) *Client {
	return &Client{
		cli: redis.NewClient(&redis.Options{
			Addr:     cfg.Addr(),
			Password: cfg.Password(),
			DB:       cfg.DB(),
		}),
	}
}

// Close closes the connection to the Redis server.
//
// It should be called when the client is no longer needed to free up resources.
// Returns an error if the operation fails.
func (c *Client) Close(_ context.Context) error {
	return c.cli.Close()
}
