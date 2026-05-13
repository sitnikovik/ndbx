package node

import "github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/node/property"

// Node represents a node in the graph.
type Node struct {
	// props holds the node properties.
	props property.Properties
	// id holds the node identifier.
	id ID
	// key holds the node key.
	key string
}

// NewNode creates a new node
// with the given id, properties and options.
func NewNode(
	id ID,
	props property.Properties,
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
func (n Node) ID() ID {
	return n.id
}

// Key returns the key of the node.
func (n Node) Key() string {
	return n.key
}

// Properties returns the properties of the node.
func (n Node) Properties() property.Properties {
	return n.props
}
