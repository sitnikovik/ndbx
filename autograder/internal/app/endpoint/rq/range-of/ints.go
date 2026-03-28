package rangeof

import (
	"net/url"
)

// Ints represents a range of integer values.
type Ints struct {
	// name is the name of the range.
	name string
	// from is the lower bound of the range.
	from Int
	// to is the upper bound of the range.
	to Int
}

// NewInts creates a new Ints instance with the given name and range of values.
func NewInts(name string, from, to Int) Ints {
	return Ints{
		name: name,
		from: from,
		to:   to,
	}
}

// URLQuery converts the range into url.Values.
//
// Example:
//
//	q := NewInts(
//	  "age",
//	  NewInt(18),
//	  NewInt(30),
//	).URLQuery().String() // "age_from=18&age_to=30"
func (r Ints) URLQuery() url.Values {
	return NewStrings(
		r.name,
		r.from.String(),
		r.to.String(),
	).URLQuery()
}
