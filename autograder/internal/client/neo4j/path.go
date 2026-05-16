package neo4j

import (
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"

	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/graph"
)

// Path executes a query and returns the path that matches the query.
func (c *Client) Path(
	ctx context.Context,
	query string,
	params map[string]any,
	key string,
) (graph.Path, error) {
	c.MustConnect()
	res, err := neo4j.ExecuteQuery(
		ctx,
		c.cli,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return graph.Path{}, err
	}
	if len(res.Records) == 0 {
		return graph.Path{}, errors.New("path not found")
	}
	p, _, err := neo4j.GetRecordValue[neo4j.Path](res.Records[0], key)
	if err != nil {
		return graph.Path{}, err
	}
	nodes := make(graph.Nodes, len(p.Nodes))
	for i, n := range p.Nodes {
		nodes[i] = graph.NewNode(
			n.GetElementId(),
			graph.PropertiesFromMap(n.GetProperties()),
		)
	}
	rels := make(graph.Relationships, len(p.Relationships))
	for i, r := range p.Relationships {
		rels[i] = graph.NewRelationship(
			r.GetElementId(),
			r.Type,
			graph.NewPoint(r.StartElementId),
			graph.NewPoint(r.EndElementId),
			graph.PropertiesFromMap(r.GetProperties()),
		)
	}
	return graph.NewPath(nodes, rels), nil
}
