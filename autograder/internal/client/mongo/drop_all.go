package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// DropAll drops all user databases in the MongoDB server.
// System databases (admin, local, config) are skipped as they cannot be dropped.
func (c *Client) DropAll(ctx context.Context) error {
	var err error
	if !c.connected {
		err = c.Connect()
		if err != nil {
			return err
		}
	}
	dbs, err := c.cli.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		return errs.Wrap(err, "failed to get all database names")
	}
	for _, db := range dbs {
		// Skip system databases that cannot be dropped
		if db == "admin" || db == "local" || db == "config" {
			continue
		}
		err = c.cli.Database(db).Drop(ctx)
		if err != nil {
			return errs.Wrap(err, "failed to drop database '%s'", db)
		}
	}
	return nil
}
