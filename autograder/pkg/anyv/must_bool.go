package anyv

// MustBool returns the raw value as a boolean value
// and panics if the type assertion did not succeed.
func (v Value) MustBool() bool {
	x, ok := v.AsBool()
	if !ok {
		panic("value is not a bool")
	}
	return x
}
