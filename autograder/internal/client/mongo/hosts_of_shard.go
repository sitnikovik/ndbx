package mongo

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"

	bsoni "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/bson"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

// HostsOfShard returns a list of hosts where the shard is running.
func (c *Client) HostsOfShard(
	ctx context.Context,
	id string,
) ([]string, error) {
	var cmd bson.M
	err := c.cli.
		Database("admin").
		RunCommand(
			ctx,
			doc.
				NewKVs(
					doc.NewKV("listShards", 1),
				).
				ToBsonD(),
		).
		Decode(&cmd)
	if err != nil {
		return nil, err
	}
	shards, ok := cmd["shards"]
	if !ok {
		return nil, errors.New("shards field not found")
	}
	items, err := bsoni.ToArray(shards)
	if err != nil {
		return nil, errors.New("shards is not an array")
	}
	for _, raw := range items {
		m, err := bsoni.ToMap(raw)
		if err != nil {
			continue
		}
		if m["_id"] != id {
			continue
		}
		v, ok := m["host"].(string)
		if !ok {
			return nil, errors.New("host is not a string")
		}
		if _, after, cut := strings.Cut(v, "/"); cut {
			return strings.Split(after, ","), nil
		}
		return []string{v}, nil
	}
	return nil, errors.New("not found")
}
