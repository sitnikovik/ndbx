package anyv

// Value defines the value that could be represented as the typed one.
type Value struct {
	// raw is a raw value to be asserted with type.
	raw any
}

func NewValue(v any) Value {
	return Value{
		raw: v,
	}
}
