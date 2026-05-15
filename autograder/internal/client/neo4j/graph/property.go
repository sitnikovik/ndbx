package graph

// Property represents a property of a node.
type Property struct {
	// value is the value of the property.
	value Value
	// name is the name of the property.
	name string
}

// NewProperty creates a new Property with the given name and value.
func NewProperty(
	name string,
	value Value,
) Property {
	return Property{
		name:  name,
		value: value,
	}
}

// Name returns the name of the property.
func (p Property) Name() string {
	return p.name
}

// Value returns the value of the property.
func (p Property) Value() Value {
	return p.value
}
