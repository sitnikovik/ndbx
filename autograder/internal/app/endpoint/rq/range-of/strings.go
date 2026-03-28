package rangeof

import (
	"net/url"
	"strings"
)

// Strings represents a range of string values.
type Strings struct {
	// name is the name of the range.
	name string
	// from is the lower bound of the range.
	from string
	// to is the upper bound of the range.
	to string
}

// NewStrings creates a new Strings instance with the given name and range of values.
func NewStrings(name, from, to string) Strings {
	return Strings{
		name: name,
		from: from,
		to:   to,
	}
}

// URLQuery converts the range into url.Values.
//
// Example:
//
//	q := NewUintRange(
//	  "date",
//	  "20230101",
//	  "20230107",
//	).URLQuery().String() // "date_from=20230101&date_to=20230107"
func (r Strings) URLQuery() url.Values {
	q := make(url.Values, 2)
	if v := strings.TrimSpace(r.from); v != "" {
		q.Set(
			r.name+"_from",
			v,
		)
	}
	if v := strings.TrimSpace(r.to); v != "" {
		q.Set(
			r.name+"_to",
			v,
		)
	}
	return q
}
