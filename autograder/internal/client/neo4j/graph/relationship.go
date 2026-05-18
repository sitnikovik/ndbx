package graph

// Relationship represents a graph relationship.
type Relationship struct {
	props Properties
	from  Point
	to    Point
	typ   string
	id    string
}

// NewRelationship creates a relationship with the provided attributes.
func NewRelationship(
	id string,
	typ string,
	from Point,
	to Point,
	props Properties,
) Relationship {
	return Relationship{
		props: props,
		from:  from,
		to:    to,
		typ:   typ,
		id:    id,
	}
}

// From returns the start point of the relationship.
func (r Relationship) From() Point {
	return r.from
}

// To returns the end point of the relationship.
func (r Relationship) To() Point {
	return r.to
}

// Type returns the relationship type.
func (r Relationship) Type() string {
	return r.typ
}

// ID returns the relationship identifier.
func (r Relationship) ID() string {
	return r.id
}

// Properties returns the relationship properties.
func (r Relationship) Properties() Properties {
	return r.props
}
