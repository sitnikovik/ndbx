package doc

import "go.mongodb.org/mongo-driver/v2/bson"

// KVs represents a slice of KV, which is a key-value pair for MongoDB operations.
type KVs []KV

// NewKVs creates a new KVs from a variadic number of KV.
func NewKVs(kvs ...KV) KVs {
	return KVs(kvs)
}

// ToBsonD converts the KVs to a bson D, which is a document used in MongoDB operations.
func (kk KVs) ToBsonD() bson.D {
	if kk.Empty() {
		return nil
	}
	bsonD := make(bson.D, 0, len(kk))
	for _, kv := range kk {
		bsonD = append(bsonD, kv.ToBsonE())
	}
	return bsonD
}

// Has checks if any of the KVs has the provided key.
func (kk KVs) Has(key string) bool {
	for _, kv := range kk {
		if kv.Has(key) {
			return true
		}
	}
	return false
}

// First returns the first KV in the KVs.
//
// If the KVs is empty, it returns an empty KV.
func (kk KVs) First() KV {
	if kk.Empty() {
		return KV{}
	}
	return kk[0]
}

// Last returns the last KV in the KVs.
//
// If the KVs is empty, it returns an empty KV.
func (kk KVs) Last() KV {
	if kk.Empty() {
		return KV{}
	}
	return kk[len(kk)-1]
}

// MustGet returns the value of the KV list if the key matches the provided key
// and panics if the key does not match.
func (kk KVs) MustGet(key string) any {
	for _, kv := range kk {
		if kv.Has(key) {
			return kv.Value()
		}
	}
	panic("key not found: " + key)
}

// Empty checks if the KVs is empty.
func (kk KVs) Empty() bool {
	return kk.Len() == 0
}

// Len returns the number of KVs in the KVs.
func (kk KVs) Len() int {
	return len(kk)
}
