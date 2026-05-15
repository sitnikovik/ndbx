package anyv

// MustBytes returns the raw value as a slice of bytes
// and panics if the type assertion not succeded.
func (v Value) MustBytes() []byte {
	x, ok := v.AsBytes()
	if !ok {
		panic("value is not a slice of bytes")
	}
	return x
}
