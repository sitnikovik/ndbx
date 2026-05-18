package anyv

import "time"

// AsTime returns the value as a time
// and a bool indicating whether the type assertion succeeded.
func (v Value) AsTime() (time.Time, bool) {
	x, ok := v.raw.(time.Time)
	return x, ok
}
