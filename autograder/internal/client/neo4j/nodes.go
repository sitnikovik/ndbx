package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/graph"
)

// QueryNodes executes a query and returns the nodes that match the query.
func (c *Client) QueryNodes(
	ctx context.Context,
	query string,
	params map[string]any,
	keys ...string,
) (graph.Nodes, error) {
	c.MustConnect()
	res, err := neo4j.ExecuteQuery(
		ctx,
		c.cli,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return graph.Nodes{}, err
	}
	nn := make([]graph.Node, 0, len(keys))
	for i, k := range keys {
		n, _, err := neo4j.GetRecordValue[neo4j.Node](res.Records[i], k)
		if err != nil {
			return graph.Nodes{}, err
		}
		nn = append(nn, graph.NewNode(
			n.GetElementId(),
			graph.PropertiesFromMap(n.GetProperties()),
			graph.WithKey(k),
		))
	}
	return graph.NewNodes(nn...), nil
}
