package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	bsonx "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/bson"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// OneBy retrieves a single document from the specified collection
// that matches the provided key-value pairs.
func (c *Client) OneBy(
	ctx context.Context,
	collection string,
	by doc.KVs,
) (doc.Document, error) {
	c.MustConnect()
	var m bson.M
	err := c.cli.
		Database(c.db).
		Collection(collection).
		FindOne(ctx, by.ToBsonD()).
		Decode(&m)
	if err != nil {
		return doc.Document{}, err
	}
	res := bsonx.
		NewBsonM(m).
		ToDocument()
	return res, nil
}
