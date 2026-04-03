package reaction

import "github.com/sitnikovik/ndbx/autograder/internal/app/creation"

// Like represents a like reaction
// for some object in the app may be reacted.
type Like struct {
	// stamp holds creation metadata.
	stamp creation.Stamp
}

// NewLike cretates a new Like instance.
func NewLike(stamp creation.Stamp) Like {
	return Like{
		stamp: stamp,
	}
}

// Created returns creation metadata.
func (l Like) Created() creation.Created {
	return l.stamp.Created()
}
