package creation

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// Created represents the creation details of some object.
type Created struct {
	// at is the time when the event was created.
	at time.Time
	// by is the user who created the event.
	by user.Identity
}

// NewCreated creates a new Created instance.
func NewCreated(
	at time.Time,
	by user.Identity,
) Created {
	return Created{
		at: at,
		by: by,
	}
}

// At returns the time when the event was created.
func (c Created) At() time.Time {
	return c.at
}

// By returns the user who created the event.
func (c Created) By() user.Identity {
	return c.by
}
