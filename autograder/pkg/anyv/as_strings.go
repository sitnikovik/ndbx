package anyv

// AsStrings returns the value as a slice of strings
// and a bool indicating whether the type assertion succeeded.
func (v Value) AsStrings() ([]string, bool) {
	x, ok := v.raw.([]string)
	return x, ok
}
