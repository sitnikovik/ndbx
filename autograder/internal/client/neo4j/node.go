package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/graph"
)

// Node executes a query and returns the node that matches the query.
func (c *Client) Node(
	ctx context.Context,
	query string,
	params map[string]any,
	key string,
) (graph.Node, error) {
	c.MustConnect()
	res, err := neo4j.ExecuteQuery(
		ctx,
		c.cli,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return graph.Node{}, err
	}
	n, _, err := neo4j.GetRecordValue[neo4j.Node](res.Records[0], key)
	if err != nil {
		return graph.Node{}, err
	}
	return graph.NewNode(
		n.GetElementId(),
		graph.PropertiesFromMap(n.GetProperties()),
		graph.WithKey(key),
	), nil
}
