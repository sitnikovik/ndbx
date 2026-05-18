package anyv

// MustString returns the raw value as a string
// and panics if the type assertion did not succeed.
func (v Value) MustString() string {
	x, ok := v.AsString()
	if !ok {
		panic("value is not a string")
	}
	return x
}
