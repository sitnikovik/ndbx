package anyv

// AsInt64 returns the value as an int64
// and a bool indicating whether the type assertion succeeded.
func (v Value) AsInt64() (int64, bool) {
	x, ok := v.raw.(int64)
	return x, ok
}
