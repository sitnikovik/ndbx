package rangeof

import (
	"net/url"
)

// UInts represents a range of unsigned integer values.
type UInts struct {
	// name is the name of the range.
	name string
	// from is the lower bound of the range.
	from UInt
	// to is the upper bound of the range.
	to UInt
}

// NewUInts creates a new UInts instance with the given name and range of values.
func NewUInts(name string, from, to UInt) UInts {
	return UInts{
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
//	  "age",
//	  NewUInt(18),
//	  NewUInt(30),
//	).URLQuery().String() // "age_from=18&age_to=30"
func (r UInts) URLQuery() url.Values {
	return NewStrings(
		r.name,
		r.from.String(),
		r.to.String(),
	).URLQuery()
}
