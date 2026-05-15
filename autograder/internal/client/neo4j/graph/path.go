package graph

type Path struct {
	nodes         Nodes
	relationships Relationships
}

func NewPath(nodes Nodes, rels Relationships) Path {
	return Path{
		nodes:         nodes,
		relationships: rels,
	}
}

func (p Path) Nodes() Nodes {
	return p.nodes
}

func (p Path) Relationships() Relationships {
	return p.relationships
}
