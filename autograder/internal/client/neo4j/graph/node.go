package graph

// Node represents a node in the graph.
type Node struct {
	// props holds the node properties.
	props Properties
	// id holds the node identifier.
	id string
	// key holds the node key.
	key string
}

// NewNode creates a new node
// with the given id, properties and options.
func NewNode(
	id string,
	props Properties,
	opts ...Option,
) Node {
	n := Node{
		id:    id,
		props: props,
	}
	for _, opt := range opts {
		opt(&n)
	}
	return n
}

// ID returns the identified of the node.
func (n Node) ID() string {
	return n.id
}

// Key returns the key of the node.
func (n Node) Key() string {
	return n.key
}

// Properties returns the properties of the node.
func (n Node) Properties() Properties {
	return n.props
}
