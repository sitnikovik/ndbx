package rangeof

import (
	"net/url"
	"time"
)

// Dates represents a range of dates.
type Dates struct {
	// name is the name of the range.
	name string
	// from is the start date in the range.
	from time.Time
	// from is the end date in the range.
	to time.Time
}

// NewDates created a new Dates instance.
func NewDates(name string, from, to time.Time) Dates {
	return Dates{
		name: name,
		from: from,
		to:   to,
	}
}

// URLQuery converts the range into url.Values.
//
// Example:
//
//	q := NewUInts(
//	  "date",
//	  time.Time{}, // started_at
//	  time.Time{}, // finished_at
//	).URLQuery().String() // "date_from=20060102&date_to=20060102"
func (r Dates) URLQuery() url.Values {
	q := make(url.Values, 2)
	if v := r.from; !v.IsZero() {
		q.Set(
			r.name+"_from",
			v.Format("20060102"),
		)
	}
	if v := r.to; !v.IsZero() {
		q.Set(
			r.name+"_to",
			v.Format("20060102"),
		)
	}
	return q
}
