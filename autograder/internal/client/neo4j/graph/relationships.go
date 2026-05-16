package graph

// Relationships represents a collection of relationships.
type Relationships []Relationship

// NewRelationships creates a new Relationships collection.
func NewRelationships(relationships ...Relationship) Relationships {
	return relationships
}

// ByType filters the collection by relationship type.
func (r Relationships) ByType(t string) Relationships {
	rels := make(Relationships, 0)
	for _, v := range r {
		if v.Type() == t {
			rels = append(rels, v)
		}
	}
	return rels
}
