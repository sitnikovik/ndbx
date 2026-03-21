package shard

import (
	"errors"
	"fmt"
)

// Shard represents a shard in MongoDB.
type Shard struct {
	// id is an identifier of the shard.
	id string
	// count is count of records stored in the shard.
	count int
	// ok defines if the shard is ok or not.
	ok bool
}

// NewShard creates a new Shard instance.
func NewShard(
	id string,
	opts ...Option,
) Shard {
	sh := Shard{
		id: id,
	}
	for _, opt := range opts {
		opt(&sh)
	}
	return sh
}

// ParseShard parses and returns a shard by the provided map
// or returns an error if not parsed.
func ParseShard(
	id string,
	m map[string]any,
) (Shard, error) {
	if id == "" {
		return Shard{}, errors.New("empty id")
	}
	if len(m) == 0 {
		return Shard{}, errors.New("empty map")
	}
	var opts []Option
	for k, v := range m {
		switch k {
		case "count":
			switch v := v.(type) {
			case int32:
				opts = append(opts, WithCount(int(v)))
			case int64:
				opts = append(opts, WithCount(int(v)))
			default:
				return Shard{},
					fmt.Errorf("expect 'count' to be type of int32 or int64 but got: %T", v)
			}
		case "ok":
			v, ok := v.(float64)
			if !ok {
				return Shard{},
					fmt.Errorf("expect 'ok' to be type of float64 but got: %T", v)
			}
			opts = append(opts, WithOk(ok))
		}
	}
	return NewShard(
		id,
		opts...,
	), nil
}

// ID returns an identifier of the shard.
func (s Shard) ID() string {
	return s.id
}

// Count returns count of records stored in the shard.
func (s Shard) Count() int {
	return s.count
}

// Ok defines if the shards is working.
func (s Shard) Ok() bool {
	return s.ok
}
