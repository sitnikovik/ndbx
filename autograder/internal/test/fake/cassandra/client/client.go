package client

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
)

// Client is a fake implementation
// of the Apache Cassandra client used for testing purposes.
type Client struct {
	// funcs is a struct that contains the functions
	// that will be used to mock the behavior
	// of the Apache Cassandra client.
	funcs funcs
}

// NewClient creates a new instance of the Client struct.
func NewClient(opts ...Option) *Client {
	fc := new(Client)
	for _, opt := range opts {
		opt(fc)
	}
	return fc
}

// Select simulates the behavior of the Select method.
func (c *Client) Select(
	ctx context.Context,
	query string,
	args ...any,
) (cassandra.Scanner, error) {
	if c.funcs.selectfn == nil {
		panic("not specified behavior for Select method")
	}
	return c.funcs.selectfn(ctx, query, args...)
}
