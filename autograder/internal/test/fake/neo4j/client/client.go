package client

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/graph"
)

// Client is a fake implementation
// of the Neo4j client used for testing purposes.
type Client struct {
	// funcs is a struct that contains the functions
	// that will be used to mock the behavior
	// of the Neo4j client.
	funcs funcs
}

// NewClient creates a new instance of the Client struct.
func NewClient(opts ...Option) *Client {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// QueryNodes simulates the behavior of the QueryNodes method.
func (c *Client) QueryNodes(
	ctx context.Context,
	query string,
	params map[string]any,
	keys ...string,
) (graph.Nodes, error) {
	if c.funcs.queryNodes == nil {
		panic("not specified behavior for QueryNodes method")
	}
	return c.funcs.queryNodes(ctx, query, params, keys...)
}
