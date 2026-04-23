package rating

import (
	"errors"
	"fmt"
	"math"
)

// number is a type constraint that allows number types for Rating.
type number interface {
	~int | ~int32 | ~int64 | ~uint8 | ~float32 | ~float64
}

// Rating represents a rating value.
type Rating float64

const (
	// None represents no rating.
	None Rating = 0
	// One represents the lowest rating.
	One Rating = 1
	// Two represents a low rating.
	Two Rating = 2
	// Three represents a middle rating.
	Three Rating = 3
	// Four represents a high rating.
	Four Rating = 4
	// Five represents the highest rating.
	Five Rating = 5
)

// NewRating creates a new Rating value.
func NewRating[T number](n T) Rating {
	return Rating(n)
}

// String returns string representation of Rating.
//
// Panics if Rating is invalid.
func (r Rating) String() string {
	if err := r.Validate(); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%.1f", r)
}

// Int returns rounded int representation of Rating.
//
// Panics if Rating is invalid.
func (r Rating) Int() int {
	return r.Round()
}

// Round returns rounded integer representation of Rating as integer.
//
// Panics if Rating is invalid.
func (r Rating) Round() int {
	if err := r.Validate(); err != nil {
		panic(err)
	}
	return int(math.Round(float64(r)))
}

// Exact returns exact value of the Rating
// that could be used in calculations.
func (r Rating) Exact() float64 {
	return float64(r)
}

// Validate returns an error if the Rating is invalid.
func (r Rating) Validate() error {
	if r.Empty() {
		return nil
	}
	if r < 1 || r > 5 {
		return errors.New("must be between 1 and 5 or equal to 0")
	}
	return nil
}

// Empty returns true if the Rating is empty or not set.
func (r Rating) Empty() bool {
	return r == None
}
