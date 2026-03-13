package session

import "time"

// Dates holds dates related to a user session.
type Dates struct {
	// createdAt is the date and time when the session was created.
	createdAt time.Time
	// updatedAt is the date and time when the session was last updated.
	updatedAt time.Time
}

// NewDates creates a new Dates instance with the given createdAt and updatedAt times.
func NewDates(createdAt, updatedAt time.Time) Dates {
	return Dates{
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

// CreatedAt returns the date and time when the session was created.
func (d Dates) CreatedAt() time.Time {
	return d.createdAt
}

// UpdatedAt returns the date and time when the session was last updated.
func (d Dates) UpdatedAt() time.Time {
	return d.updatedAt
}
