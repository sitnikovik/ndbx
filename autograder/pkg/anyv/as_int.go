package anyv

// AsInt returns the value as an int
// and a bool indicating whether the type assertion succeeded.
func (v Value) AsInt() (int, bool) {
	x, ok := v.raw.(int)
	return x, ok
}
