package property

// Properties represents a collection of Property.
type Properties struct {
	// m is the underlying map of property name to Property.
	m map[string]Property
}

// NewProperties creates a new Properties from a variadic list of Property.
func NewProperties(pp ...Property) Properties {
	m := make(map[string]Property, len(pp))
	for _, p := range pp {
		m[p.name] = p
	}
	return Properties{
		m: m,
	}
}

// PropertiesFromMap creates a Properties from a map of string to any.
func PropertiesFromMap(m map[string]any) Properties {
	pm := make(map[string]Property, len(m))
	for k, v := range m {
		pm[k] = NewProperty(k, NewValue(v))
	}
	return Properties{
		m: pm,
	}
}

// ByName returns the property with the given name, if it exists.
func (p Properties) ByName(name string) Property {
	return p.m[name]
}

// Len returns the number of properties.
func (p Properties) Len() int {
	return len(p.m)
}
