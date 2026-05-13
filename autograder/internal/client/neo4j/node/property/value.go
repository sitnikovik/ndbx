package property

// Value represents the value of a property.
type Value struct {
	// raw is the raw value of the property.
	raw any
}

// NewValue creates a new Value from the raw value.
func NewValue(raw any) Value {
	return Value{
		raw: raw,
	}
}

// Raw returns the raw value of the property.
func (v Value) Raw() any {
	return v.raw
}

// MustInt returns the raw value as an int.
func (v Value) MustInt() int {
	x, ok := v.raw.(int)
	if !ok {
		panic("value is not an int")
	}
	return x
}

// MustString returns the raw value as a string.
func (v Value) MustString() string {
	x, ok := v.raw.(string)
	if !ok {
		panic("value is not a string")
	}
	return x
}
