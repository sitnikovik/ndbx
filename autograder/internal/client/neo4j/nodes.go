package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/node"
	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/node/property"
)

// Nodes executes a query and returns the nodes that match the query.
func (c *Client) QueryNodes(
	ctx context.Context,
	query string,
	params map[string]any,
	keys ...string,
) (node.Nodes, error) {
	c.MustConnect()
	res, err := neo4j.ExecuteQuery(
		ctx,
		c.cli,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return node.Nodes{}, err
	}
	nn := make([]node.Node, 0, len(keys))
	for i, k := range keys {
		n, _, err := neo4j.GetRecordValue[neo4j.Node](res.Records[i], k)
		if err != nil {
			return node.Nodes{}, err
		}
		nn = append(nn, node.NewNode(
			node.NewID(n.GetElementId()),
			property.PropertiesFromMap(n.GetProperties()),
		))
	}
	return node.NewNodes(nn...), nil
}
