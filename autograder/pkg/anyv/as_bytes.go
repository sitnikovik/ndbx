package anyv

// AsBytes returns the value as a slice of bytes
// and a bool indicating whether the type assertion succeeded.
func (v Value) AsBytes() ([]byte, bool) {
	x, ok := v.raw.([]byte)
	return x, ok
}
