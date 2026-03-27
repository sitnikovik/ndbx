package shard

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"

	bsoni "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/bson"
)

// Shards represents a list of shards in MongoDB.
type Shards []Shard

// NewShards creates a new Shards instance.
func NewShards(sh ...Shard) Shards {
	return Shards(sh)
}

// ParseShards parses shards from a map or bson value.
// Accepts map[string]map[string]any, bson.M, or bson.D.
func ParseShards(v any) (Shards, error) {
	var mm map[string]map[string]any
	switch val := v.(type) {
	case map[string]map[string]any:
		mm = val
	case bson.M, bson.D:
		entries, err := bsoni.ToMap(val)
		if err != nil {
			return nil, errors.New("shards is not a map")
		}
		mm = make(map[string]map[string]any, len(entries))
		for id, v := range entries {
			m, err := bsoni.ToMap(v)
			if err != nil {
				return nil, fmt.Errorf("shard '%s': %w", id, err)
			}
			mm[id] = m
		}
	default:
		return nil, fmt.Errorf("unexpected type %T", v)
	}
	if len(mm) == 0 {
		return nil, errors.New("empty map")
	}
	shh := make(Shards, 0, len(mm))
	for id, m := range mm {
		sh, err := ParseShard(id, m)
		if err != nil {
			return nil, fmt.Errorf("error for '%s': %w", id, err)
		}
		shh = append(shh, sh)
	}
	return shh, nil
}
