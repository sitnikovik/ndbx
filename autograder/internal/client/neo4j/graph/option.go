package graph

// Option represents the functional option
// to modify a Node on its creation.
type Option func(*Node)

// WithKey sets the key of the Node.
func WithKey(key string) Option {
	return func(n *Node) {
		n.key = key
	}
}
