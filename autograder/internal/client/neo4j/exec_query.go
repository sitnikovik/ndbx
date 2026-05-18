package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// ExecQuery executes the given query with the given parameters
// against the Neo4j database and returns an error if the query fails.
func (c *Client) ExecQuery(
	ctx context.Context,
	query string,
	params map[string]any,
) error {
	c.MustConnect()
	_, err := neo4j.ExecuteQuery(
		ctx,
		c.cli,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	return err
}
