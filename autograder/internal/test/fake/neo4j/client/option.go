package client

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/graph"
)

// Option defines a functional option for configuring the fake Client.
type Option func(*Client)

// WithQueryNodes sets the QueryNodes function for the fake Client.
func WithQueryNodes(
	fn func(
		ctx context.Context,
		query string,
		params map[string]any,
		keys ...string,
	) (graph.Nodes, error),
) Option {
	return func(c *Client) {
		c.funcs.queryNodes = fn
	}
}
