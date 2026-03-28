package rangeof

import "strconv"

// intval is an interface for all integer types.
type intval interface {
	int | int8 | int16 | int32 | int64
}

// Int represents an integer value for the range of numbers.
type Int struct {
	// v is the integer value.
	v int64
	// force indicates whether to include the value
	// even if it's zero.
	force bool
}

// NewInt creates a new Int instance.
func NewInt[T intval](v T) Int {
	return Int{
		v: int64(v),
	}
}

// NewAnyInt creates a new Int instance
// with any integer type even if it's zero.
func NewAnyInt[T intval](v T) Int {
	return Int{
		v:     int64(v),
		force: true,
	}
}

// String returns string representation of the integer.
func (n Int) String() string {
	if n.v == 0 && !n.force {
		return ""
	}
	return strconv.FormatInt(n.v, 10)
}
