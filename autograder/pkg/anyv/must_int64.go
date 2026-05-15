package anyv

// MustInt64 returns the raw value as an int65
// and panics if the type assertion not succeded.
func (v Value) MustInt64() int64 {
	x, ok := v.AsInt64()
	if !ok {
		panic("value is not an int64")
	}
	return x
}
