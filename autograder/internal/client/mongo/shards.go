package mongo

import (
	"context"
	"errors"
	"fmt"

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
	raw, ok := v.(bson.M)
	if !ok {
		return nil, errors.New("shards is not a map")
	}
	mm := make(map[string]map[string]any, len(raw))
	for id, val := range raw {
		shardMap, ok := val.(bson.M)
		if !ok {
			return nil,
				fmt.Errorf(
					"shard '%s' has unexpected type %T",
					id,
					val,
				)
		}
		mm[id] = map[string]any(shardMap)
	}
	return shard.ParseShards(mm)
}
