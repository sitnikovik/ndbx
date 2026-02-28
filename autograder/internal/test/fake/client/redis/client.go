package redis

import (
	"context"
	"time"
)

// FakeClient is a mock implementation of a Redis client for testing purposes.
//
// It allows you to define custom behavior for Redis commands by providing functional options.
type FakeClient struct {
	// funcs holds the set of functions that can be configured to simulate Redis command behavior.
	funcs funcs
}

// NewFakeClient creates a new instance of FakeClient with the provided options.
func NewFakeClient(opts ...Option) *FakeClient {
	c := &FakeClient{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Ping simulates a Redis PING command, returning an error if the client is not healthy.
//
// Panics if the behavior for the Ping method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) Ping(ctx context.Context) error {
	if c.funcs.ping == nil {
		panic("not specified behavior for Ping method")
	}
	return c.funcs.ping(ctx)
}

// FlushAll simulates a Redis FLUSHALL command, returning an error if the operation fails.
//
// Panics if the behavior for the FlushAll method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) FlushAll(ctx context.Context) error {
	if c.funcs.flushAll == nil {
		panic("not specified behavior for FlushAll method")
	}
	return c.funcs.flushAll(ctx)
}

// Empty simulates a check to determine if the Redis database is empty,
// returning a boolean indicating emptiness or an error if the operation fails.
//
// Panics if the behavior for the Empty method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) Empty(ctx context.Context) (bool, error) {
	if c.funcs.empty == nil {
		panic("not specified behavior for Empty method")
	}
	return c.funcs.empty(ctx)
}

// Get simulates a Redis GET command, returning the value associated with the given key
// or an error if the operation fails.
//
// Panics if the behavior for the Get method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) Get(
	ctx context.Context,
	key string,
) (string, error) {
	if c.funcs.get == nil {
		panic("not specified behavior for Get method")
	}
	return c.funcs.get(ctx, key)
}

// Set simulates a Redis SET command, setting the value for a given key with an optional TTL,
// and returning an error if the operation fails.
//
// Panics if the behavior for the Set method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) Set(
	ctx context.Context,
	key string,
	value string,
	ttl time.Duration,
) error {
	if c.funcs.set == nil {
		panic("not specified behavior for Set method")
	}
	return c.funcs.set(ctx, key, value, ttl)
}

// Has simulates a Redis EXISTS command, checking if a given key exists in the database,
// and returning a boolean indicating existence or an error if the operation fails.
//
// Panics if the behavior for the Has method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) Has(
	ctx context.Context,
	key string,
) (bool, error) {
	if c.funcs.has == nil {
		panic("not specified behavior for Has method")
	}
	return c.funcs.has(ctx, key)
}

// HGet simulates a Redis HGET command, retrieving the value of a field in a hash stored at a key,
// and returning the value or an error if the operation fails.
//
// Panics if the behavior for the HGet method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) HGet(
	ctx context.Context,
	key string,
	field string,
) (string, error) {
	if c.funcs.hget == nil {
		panic("not specified behavior for HGet method")
	}
	return c.funcs.hget(ctx, key, field)
}

// HGetAll simulates a Redis HGETALL command, retrieving all fields and values of a hash stored at a key,
// and returning a map of field-value pairs or an error if the operation fails.
//
// Panics if the behavior for the HGetAll method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) HGetAll(
	ctx context.Context,
	key string,
) (map[string]string, error) {
	if c.funcs.hgetall == nil {
		panic("not specified behavior for HGetAll method")
	}
	return c.funcs.hgetall(ctx, key)
}

// TTL simulates a Redis TTL command, retrieving the time-to-live of a key,
// and returning the duration or an error if the operation fails.
//
// Panics if the behavior for the TTL method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) TTL(
	ctx context.Context,
	key string,
) (time.Duration, error) {
	if c.funcs.ttl == nil {
		panic("not specified behavior for TTL method")
	}
	return c.funcs.ttl(ctx, key)
}

// Type simulates a Redis TYPE command, retrieving the data type of the value stored at a key,
// and returning the type as a string or an error if the operation fails.
//
// Panics if the behavior for the Type method is not specified, as it is expected to be defined for testing purposes.
func (c *FakeClient) Type(
	ctx context.Context,
	key string,
) (string, error) {
	if c.funcs.typ == nil {
		panic("not specified behavior for Type method")
	}
	return c.funcs.typ(ctx, key)
}

// Close simulates closing the connection to the Redis server.
//
// It should be called when the client is no longer needed to free up resources.
// Returns an error if the operation fails.
func (c *FakeClient) Close(ctx context.Context) error {
	if c.funcs.close != nil {
		return c.funcs.close(ctx)
	}
	return nil
}
