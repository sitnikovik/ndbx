package doc

// Indexes represents a collection of MongoDB indexes.
type Indexes []Index

// NewIndexes creates a new Indexes with the provided indexes.
func NewIndexes(indexes ...Index) Indexes {
	return indexes
}

// HasAllFor checks if any of the indexes in the Indexes has all of the provided keys.
func (ii Indexes) HasAllFor(keys ...string) bool {
	for _, idx := range ii {
		if idx.HasAllFor(keys...) {
			return true
		}
	}
	return false
}

// HasAnyOf checks if any of the indexes in the Indexes has any of the provided keys.
func (ii Indexes) HasAnyOf(keys ...string) bool {
	for _, idx := range ii {
		if idx.HasAnyOf(keys...) {
			return true
		}
	}
	return false
}

// For returns the first index in the Indexes that has all of the provided keys.
func (ii Indexes) For(keys ...string) Index {
	for _, idx := range ii {
		if idx.HasAllFor(keys...) {
			return idx
		}
	}
	return Index{}
}
