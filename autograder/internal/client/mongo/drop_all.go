package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// DropAll truncates all collections in all user databases in the MongoDB server.
// System databases (admin, local, config) are skipped.
// This is similar to TRUNCATE in SQL - it removes all documents but keeps the database structure.
func (c *Client) DropAll(ctx context.Context) error {
	c.MustConnect()
	dbs, err := c.cli.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		return errs.Wrap(
			err,
			"failed to get all database names",
		)
	}
	for _, db := range dbs {
		if db == "admin" || db == "local" || db == "config" {
			continue
		}
		database := c.cli.Database(db)
		collections, err := database.ListCollectionNames(ctx, bson.D{})
		if err != nil {
			return errs.Wrap(
				err,
				"failed to list collections in database '%s'",
				db,
			)
		}
		for _, collName := range collections {
			_, err = database.Collection(collName).DeleteMany(ctx, bson.D{})
			if err != nil {
				return errs.Wrap(
					err,
					"failed to truncate collection '%s' in database '%s'",
					collName,
					db,
				)
			}
		}
	}
	return nil
}
