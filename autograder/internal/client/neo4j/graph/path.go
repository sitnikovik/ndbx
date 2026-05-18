package graph

// Path represents a graph path with ordered nodes and relationships.
type Path struct {
	nodes         Nodes
	relationships Relationships
}

// NewPath creates a path from the provided nodes and relationships.
func NewPath(nodes Nodes, rels Relationships) Path {
	return Path{
		nodes:         nodes,
		relationships: rels,
	}
}

// Nodes returns the nodes that belong to the path.
func (p Path) Nodes() Nodes {
	return p.nodes
}

// Relationships returns the relationships that belong to the path.
func (p Path) Relationships() Relationships {
	return p.relationships
}
