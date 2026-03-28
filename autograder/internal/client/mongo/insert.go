package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Insert inserts the list of documents into the specified collection.
// Returns a slice of inserted document IDs and an error if any.
func (c *Client) Insert(
	ctx context.Context,
	collection string,
	kvs ...doc.KVs,
) ([]string, error) {
	if len(kvs) == 0 {
		return nil, errs.Wrap(errs.ErrInvalidParam, "empty key-value documents")
	}
	c.MustConnect()
	lst := make([]bson.D, len(kvs))
	for i, kv := range kvs {
		lst[i] = kv.ToBsonD()
	}
	result, err := c.cli.
		Database(c.db).
		Collection(collection).
		InsertMany(ctx, lst)
	if err != nil {
		return nil, err
	}
	ids := make([]string, len(result.InsertedIDs))
	for i, id := range result.InsertedIDs {
		ids[i] = id.(string)
	}
	return ids, nil
}
