package anyv

// IsNil reports whether the wrapped value is nil.
func (v Value) IsNil() bool {
	return v.raw == nil
}
