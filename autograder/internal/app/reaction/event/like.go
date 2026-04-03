package event

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	common "github.com/sitnikovik/ndbx/autograder/internal/app/reaction"
)

// Like represents a like reaction
// for some object in the app may be reacted.
type Like struct {
	// orig is Like instance.
	orig common.Like
	// ev is metadata describes what the event has been liked.
	ev Event
}

// NewLike cretates a new Like instance.
func NewLike(orig common.Like, ev Event) Like {
	return Like{
		orig: orig,
		ev:   ev,
	}
}

// Created returns creation metadata.
func (l Like) Created() creation.Created {
	return l.orig.Created()
}
