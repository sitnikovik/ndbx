package graph

// Nodes represents a collection of nodes.
type Nodes []Node

// NewNodes creates a new Nodes collection from the given nodes.
func NewNodes(nodes ...Node) Nodes {
	return nodes
}

// ByKey returns the node with the given key.
func (n Nodes) ByKey(key string) Node {
	for _, v := range n {
		if v.Key() == key {
			return v
		}
	}
	return Node{}
}
