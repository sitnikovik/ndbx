package client

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/graph"
)

// funcs holds function fields
// for mocking the behavior of the Client methods.
type funcs struct {
	// QueryNodes is a function that will be used
	// to mock the behavior of the QueryNodes method.
	queryNodes func(
		ctx context.Context,
		query string,
		params map[string]any,
		keys ...string,
	) (graph.Nodes, error)
	// deleteAll is a function that will be used
	// to mock the behavior of the DeleteAll method.
	deleteAll func(ctx context.Context) error
}
