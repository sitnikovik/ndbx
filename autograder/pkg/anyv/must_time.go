package anyv

import "time"

// MustTime returns the raw value as a time
// and panics if the type assertion not succeded.
func (v Value) MustTime() time.Time {
	x, ok := v.AsTime()
	if !ok {
		panic("value is not a time")
	}
	return x
}
