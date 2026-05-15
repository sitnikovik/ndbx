package graph

type Relationship struct {
	props Properties
	from  Point
	to    Point
	typ   string
	id    string
}

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

func (r Relationship) From() Point {
	return r.from
}

func (r Relationship) To() Point {
	return r.to
}

func (r Relationship) Type() string {
	return r.typ
}

func (r Relationship) ID() string {
	return r.id
}

func (r Relationship) Properties() Properties {
	return r.props
}
