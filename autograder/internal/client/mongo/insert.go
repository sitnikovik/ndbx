package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Insert inserts the list of documents into the specified collection.
func (c *Client) Insert(
	ctx context.Context,
	collection string,
	kvs ...doc.KVs,
) error {
	if len(kvs) == 0 {
		return errs.Wrap(errs.ErrInvalidParam, "empty key-value documents")
	}
	c.MustConnect()
	lst := make([]bson.D, len(kvs))
	for i, kv := range kvs {
		lst[i] = kv.ToBsonD()
	}
	_, err := c.cli.
		Database(c.db).
		Collection(collection).
		InsertMany(ctx, lst)
	return err
}
