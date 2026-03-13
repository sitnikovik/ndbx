package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	bsonx "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/bson"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// AllBy retrieves all documents from the specified collection
// that match the provided key-value pairs.
func (c *Client) AllBy(
	ctx context.Context,
	collection string,
	by doc.KVs,
) (doc.Documents, error) {
	c.MustConnect()
	cur, err := c.cli.
		Database(c.db).
		Collection(collection).
		Find(ctx, by.ToBsonD())
	if err != nil {
		return nil, err
	}
	var mm []bson.M
	err = cur.All(ctx, &mm)
	if err != nil {
		return nil, err
	}
	res := bsonx.
		NewBsonMList(mm...).
		ToDocuments()
	return res, nil
}
