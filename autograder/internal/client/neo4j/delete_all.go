package neo4j

import "context"

// DeleteAll removes all nodes and relationships from the current Neo4j database
// while preserving the schema, including indexes and constraints.
func (c *Client) DeleteAll(ctx context.Context) error {
	return c.ExecQuery(
		ctx,
		"MATCH (n) DETACH DELETE n",
		nil,
	)
}
