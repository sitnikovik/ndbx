package filter

import "github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction/enum/like"

// Like represents a like filter.
type Like struct {
	// v holds the value of the filter.
	v like.Value
	// set defines whether the filter is set or not.
	set bool
}

// NewLike creates a new Like filter.
func NewLike(v like.Value) Like {
	return Like{
		v:   v,
		set: true,
	}
}

// Value returns the value of the filter.
func (l Like) Value() like.Value {
	return l.v
}

// Empty defines whether the filter is set or not.
func (l Like) Empty() bool {
	return !l.set
}
