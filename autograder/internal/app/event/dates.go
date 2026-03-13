package event

import "time"

// Dates represents the start and end times of an event.
type Dates struct {
	// startedAt is the time when the event starts.
	startedAt time.Time
	// finishedAt is the time when the event finishes.
	finishedAt time.Time
}

// NewDates creates a new Dates instance with the provided start and end times.
func NewDates(startedAt, finishedAt time.Time) Dates {
	return Dates{
		startedAt:  startedAt,
		finishedAt: finishedAt,
	}
}

// StartedAt returns the start time of the event.
func (d Dates) StartedAt() time.Time {
	return d.startedAt
}

// FinishedAt returns the end time of the event.
func (d Dates) FinishedAt() time.Time {
	return d.finishedAt
}
