package cassandra

import (
	"context"
	"fmt"
)

// TruncateKeyspace truncates all tables in the keyspace, effectively removing all data
// while preserving the schema (tables, types, indexes).
//
// It queries the system schema to list all tables in the keyspace,
// then executes a TRUNCATE command on each one.
//
// Returns an error if listing tables fails or if any table cannot be truncated.
func (c *Client) TruncateKeyspace(ctx context.Context) error {
	c.MustConnect()
	keyspace := c.cfg.Database().Keyspace()
	iter := c.cluster.
		Query(
			`SELECT table_name FROM system_schema.tables WHERE keyspace_name = ?`,
			keyspace,
		).
		Iter()
	var tableName string
	for iter.Scan(&tableName) {
		err := c.ExecQuery(
			ctx,
			fmt.Sprintf(
				`TRUNCATE "%s"."%s"`,
				keyspace,
				tableName,
			),
		)
		if err != nil {
			return fmt.Errorf(
				"failed to truncate table %s: %w",
				tableName,
				err,
			)
		}
	}

	return iter.Close()
}
