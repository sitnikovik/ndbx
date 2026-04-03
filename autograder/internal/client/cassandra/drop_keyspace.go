package cassandra

import "context"

// DropKeyspace drops the keyspace in all nodes of the cluster
// and returns an error if any of the nodes failed.
func (c *Client) DropKeyspace(ctx context.Context) error {
	return c.ExecQuery(
		ctx,
		"DROP KEYSPACE IF EXISTS "+c.cfg.Database().Keyspace(),
	)
}
