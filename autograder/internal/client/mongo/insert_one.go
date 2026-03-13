package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// InsertOne inserts a single document into the specified collection.
func (c *Client) InsertOne(
	ctx context.Context,
	collection string,
	kvs doc.KVs,
) error {
	c.MustConnect()
	_, err := c.cli.
		Database(c.db).
		Collection(collection).
		InsertOne(ctx, kvs.ToBsonD())
	return err
}
