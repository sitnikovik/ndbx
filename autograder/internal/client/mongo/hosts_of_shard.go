package mongo

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"

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
	raw, _ := cmd["shards"].([]any)
	for _, raw := range raw {
		m, ok := raw.(map[string]any)
		if !ok {
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
