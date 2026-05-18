package anyv

// AsBool returns the value as a boolean
// and a bool indicating whether the type assertion succeeded.
func (v Value) AsBool() (bool, bool) {
	x, ok := v.raw.(bool)
	return x, ok
}
