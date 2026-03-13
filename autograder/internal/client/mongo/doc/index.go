package doc

import (
	"slices"
)

// Index represents a MongoDB index, which is defined by a set of keys.
type Index struct {
	// Keys is a slice of key-value pairs representing the fields to be indexed.
	keys []string
	// unique indicates whether the index is unique or not.
	unique bool
}

// NewIndex creates a new Index with the provided keys.
func NewIndex(keys ...string) Index {
	return Index{
		keys: keys,
	}
}

// NewUniqueIndex creates a new unique Index with the provided keys.
func NewUniqueIndex(keys ...string) Index {
	return Index{
		keys:   keys,
		unique: true,
	}
}

// HasAllFor checks if the index has exactly all of the provided keys.
func (i Index) HasAllFor(keys ...string) bool {
	if len(i.keys) != len(keys) {
		return false
	}
	for i, k := range i.keys {
		if k != keys[i] {
			return false
		}
	}
	return true
}

// HasAnyOf checks if the index has any of the provided keys.
func (i Index) HasAnyOf(keys ...string) bool {
	for _, k := range i.keys {
		if slices.Contains(keys, k) {
			return true
		}
	}
	return false
}

// Unique returns true if the index is unique, false otherwise.
func (i Index) Unique() bool {
	return i.unique
}

// Empty returns true if the index has no keys, false otherwise.
func (i Index) Empty() bool {
	return len(i.keys) == 0
}
