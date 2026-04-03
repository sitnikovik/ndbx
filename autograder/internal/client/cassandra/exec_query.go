package cassandra

import "context"

// ExecQuery executes the given query and returns an error if any.
func (c *Client) ExecQuery(
	_ context.Context,
	query string,
) error {
	c.MustConnect()
	return c.cluster.
		Query(query).
		Exec()
}
