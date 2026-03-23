package rangeof

import "strconv"

// uintval is an interface for all unsigned integer types.
type uintval interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// UInt represents an unsigned integer value for the range of numbers.
type UInt struct {
	// v is the unsigned integer value.
	v uint64
	// force indicates whether to include the value
	// even if it's zero.
	force bool
}

// NewUInt creates a new UInt instance.
func NewUInt[T uintval](v T) UInt {
	return UInt{
		v: uint64(v),
	}
}

// NewAnyUInt creates a new UInt instance
// with any unsigned integer type even if it's zero.
func NewAnyUInt[T uintval](v T) UInt {
	return UInt{
		v:     uint64(v),
		force: true,
	}
}

// String returns string representation of the integer.
func (n UInt) String() string {
	if n.v == 0 && !n.force {
		return ""
	}
	return strconv.FormatUint(n.v, 10)
}
