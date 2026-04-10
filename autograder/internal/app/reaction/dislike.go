package reaction

import "github.com/sitnikovik/ndbx/autograder/internal/app/creation"

// Dislike represents a dislike reaction
// for some object in the app may be reacted.
type Dislike struct {
	// stamp holds creation metadata.
	stamp creation.Stamp
}

// NewDislike cretates a new Dislike instance.
func NewDislike(stamp creation.Stamp) Dislike {
	return Dislike{
		stamp: stamp,
	}
}

// Created returns creation metadata.
func (l Dislike) Created() creation.Created {
	return l.stamp.Created()
}
