package anyv

// MustInt returns the raw value as an int
// and panics if the type assertion not succeded.
func (v Value) MustInt() int {
	x, ok := v.AsInt()
	if !ok {
		panic("value is not an int")
	}
	return x
}
