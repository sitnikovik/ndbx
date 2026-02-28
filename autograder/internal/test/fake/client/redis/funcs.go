package redis

import (
	"context"
	"time"
)

// funcs defines the set of functions that can be configured for the FakeClient.
type funcs struct {
	// ping simulates a Redis PING command.
	ping func(ctx context.Context) error
	// flushAll simulates a Redis FLUSHALL command.
	flushAll func(ctx context.Context) error
	// empty simulates a check to determine if the Redis database is empty.
	empty func(ctx context.Context) (bool, error)
	// get simulates a Redis GET command.
	get func(
		ctx context.Context,
		key string,
	) (string, error)
	// set simulates a Redis SET command.
	set func(
		ctx context.Context,
		key string,
		value string,
		ttl time.Duration,
	) error
	// has simulates a Redis EXISTS command.
	has func(
		ctx context.Context,
		key string,
	) (bool, error)
	// hget simulates a Redis HGET command.
	hget func(
		ctx context.Context,
		key string,
		field string,
	) (string, error)
	// hgetall simulates a Redis HGETALL command.
	hgetall func(
		ctx context.Context,
		key string,
	) (map[string]string, error)
	// ttl simulates a Redis TTL command.
	ttl func(
		ctx context.Context,
		key string,
	) (time.Duration, error)
	// type simulates a Redis TYPE command.
	typ func(
		ctx context.Context,
		key string,
	) (string, error)
	// close simulates closing the connection to the Redis server.
	close func(ctx context.Context) error
}
