package redis

import (
	"context"
	"time"
)

// Option defines a functional option for configuring the FakeClient.
type Option func(*FakeClient)

// WithPing sets the function that will be executed
// when the FakeClient's Ping method is called.
func WithPing(
	fn func(ctx context.Context) error,
) Option {
	return func(c *FakeClient) {
		c.funcs.ping = fn
	}
}

// WithFlushAll sets the function that will be executed
// when the FakeClient's FlushAll method is called.
func WithFlushAll(
	fn func(ctx context.Context) error,
) Option {
	return func(c *FakeClient) {
		c.funcs.flushAll = fn
	}
}

// WithEmpty sets the function that will be executed
// when the FakeClient's Empty method is called.
func WithEmpty(
	fn func(ctx context.Context) (bool, error),
) Option {
	return func(c *FakeClient) {
		c.funcs.empty = fn
	}
}

// WithGet sets the function that will be executed
// when the FakeClient's Get method is called.
func WithGet(
	fn func(
		ctx context.Context,
		key string,
	) (string, error),
) Option {
	return func(c *FakeClient) {
		c.funcs.get = fn
	}
}

// WithSet sets the function that will be executed
// when the FakeClient's Set method is called.
func WithSet(
	fn func(
		ctx context.Context,
		key string,
		value string, ttl time.Duration,
	) error,
) Option {
	return func(c *FakeClient) {
		c.funcs.set = fn
	}
}

// WithHas sets the function that will be executed
// when the FakeClient's Has method is called.
func WithHas(
	fn func(
		ctx context.Context,
		key string,
	) (bool, error),
) Option {
	return func(c *FakeClient) {
		c.funcs.has = fn
	}
}

// WithHGet sets the function that will be executed
// when the FakeClient's HGet method is called.
func WithHGet(
	fn func(
		ctx context.Context,
		key string,
		field string,
	) (string, error),
) Option {
	return func(c *FakeClient) {
		c.funcs.hget = fn
	}
}

// WithHGetAll sets the function that will be executed
// when the FakeClient's HGetAll method is called.
func WithHGetAll(
	fn func(
		ctx context.Context,
		key string,
	) (map[string]string, error),
) Option {
	return func(c *FakeClient) {
		c.funcs.hgetall = fn
	}
}

// WithTTL sets the function that will be executed
// when the FakeClient's TTL method is called.
func WithTTL(
	fn func(
		ctx context.Context,
		key string,
	) (time.Duration, error),
) Option {
	return func(c *FakeClient) {
		c.funcs.ttl = fn
	}
}

// WithType sets the function that will be executed
// when the FakeClient's Type method is called.
func WithType(
	fn func(
		ctx context.Context,
		key string,
	) (string, error),
) Option {
	return func(c *FakeClient) {
		c.funcs.typ = fn
	}
}

// WithClose sets the function that will be executed
// when the FakeClient's Close method is called.
func WithClose(
	fn func(ctx context.Context) error,
) Option {
	return func(c *FakeClient) {
		c.funcs.close = fn
	}
}
