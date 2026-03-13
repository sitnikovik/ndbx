package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	bsonx "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/bson"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// Indexes retrieves the list of indexes for the specified collection.
func (c *Client) Indexes(
	ctx context.Context,
	collection string,
) (doc.Indexes, error) {
	c.MustConnect()
	ind := c.cli.
		Database(c.db).
		Collection(collection).
		Indexes()
	cur, err := ind.List(ctx)
	if err != nil {
		return nil, err
	}
	var mm []bson.M
	err = cur.All(ctx, &mm)
	if err != nil {
		return nil, err
	}
	idxx := bsonx.
		NewBsonMList(mm...).
		ToIndexes()
	return idxx, nil
}
