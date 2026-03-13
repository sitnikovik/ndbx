package doc

import "go.mongodb.org/mongo-driver/v2/bson"

// KV represents a key-value pair for MongoDB operations.
type KV struct {
	// k is the key of the key-value pair.
	k string
	// v is the value of the key-value pair.
	v any
}

// NewKV creates a new key-value pair for MongoDB operations.
func NewKV(key string, value any) KV {
	return KV{
		k: key,
		v: value,
	}
}

// Key returns the key of the key-value pair.
func (k KV) Key() string {
	return k.k
}

// Value returns the value of the key-value pair.
func (k KV) Value() any {
	return k.v
}

// ToBsonE converts the KV to a bson E, which is a key-value pair used in MongoDB operations.
func (k KV) ToBsonE() bson.E {
	return bson.E{
		Key:   k.Key(),
		Value: k.Value(),
	}
}

// Has checks if the key of the KV is equal to the provided key.
func (k KV) Has(key string) bool {
	return k.Key() == key
}

// MustGet returns the value of the KV if the key matches the provided key
// and panics if the key does not match.
func (k KV) MustGet(key string) any {
	if !k.Has(key) {
		panic("key not found: " + key)
	}
	return k.Value()
}
