package node

// Nodes represents a collection of nodes.
type Nodes struct {
	// list is the list of nodes.
	list []Node
}

// NewNodes creates a new Nodes collection from the given nodes.
func NewNodes(nodes ...Node) Nodes {
	return Nodes{
		list: nodes,
	}
}

// ByKey returns the node with the given key.
func (n Nodes) ByKey(key string) Node {
	for _, v := range n.list {
		if v.Key() == key {
			return v
		}
	}
	return Node{}
}

// Len returns the number of nodes.
func (n Nodes) Len() int {
	return len(n.list)
}

// Empty returns true if the collection is empty.
func (n Nodes) Empty() bool {
	return n.Len() == 0
}
