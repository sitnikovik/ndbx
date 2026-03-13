package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// ByID retrieves a single document
// from the specified collectionthat matches the given ID.
func (c *Client) ByID(
	ctx context.Context,
	collection string,
	id string,
) (doc.Document, error) {
	objid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return doc.Document{}, err
	}
	return c.OneBy(
		ctx,
		collection,
		doc.NewKVs(
			doc.NewKV(
				"_id",
				objid,
			),
		),
	)
}
