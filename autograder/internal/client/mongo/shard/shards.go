package shard

import (
	"errors"
	"fmt"
)

// Shards represents a list of shards in MongoDB.
type Shards []Shard

// NewShards creates a new Shards instance.
func NewShards(sh ...Shard) Shards {
	return Shards(sh)
}

// ParseShards parses and returns a list of shards
// by the provided map or returns an error if not parsed.
func ParseShards(mm map[string]map[string]any) (Shards, error) {
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
