package anyv

// MustStrings returns the raw value as a slice of strings
// and panics if the type assertion not succeded.
func (v Value) MustStrings() []string {
	x, ok := v.AsStrings()
	if !ok {
		panic("value is not a slice of strings")
	}
	return x
}
