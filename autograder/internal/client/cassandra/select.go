package cassandra

import (
	"context"

	"github.com/gocql/gocql"
)

// Select selects the rows from the given query
// and returns an iterator to scan the rows.
func (c *Client) Select(
	_ context.Context,
	query string,
	args ...any,
) (*gocql.Iter, error) {
	c.MustConnect()
	return c.cluster.
		Query(query, args...).
		Iter(), nil
}
