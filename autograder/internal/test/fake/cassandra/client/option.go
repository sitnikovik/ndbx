package client

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
)

// Option defines a functional option for configuring the fake Client.
type Option func(*Client)

// WithSelect sets the select function for the fake Client.
func WithSelect(
	fn func(
		ctx context.Context,
		query string,
		args ...any,
	) (cassandra.Scanner, error),
) Option {
	return func(c *Client) {
		c.funcs.selectfn = fn
	}
}
