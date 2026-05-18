package anyv

// AsString returns the value as a string
// and a bool indicating whether the type assertion succeeded.
func (v Value) AsString() (string, bool) {
	x, ok := v.raw.(string)
	return x, ok
}
