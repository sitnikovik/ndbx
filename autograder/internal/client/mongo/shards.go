package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/shard"
)

// Shards retrieves the shard
// information for the specified collection.
func (c *Client) Shards(
	ctx context.Context,
	collection string,
) (shard.Shards, error) {
	c.MustConnect()
	var cmd bson.M
	err := c.cli.
		Database(c.db).
		RunCommand(
			ctx,
			doc.
				NewKVs(
					doc.NewKV("collStats", collection),
				).
				ToBsonD(),
		).
		Decode(&cmd)
	if err != nil {
		return nil, err
	}
	v, ok := cmd["shards"]
	if !ok {
		return nil, errors.New("not found")
	}
	return shard.ParseShards(v)
}
