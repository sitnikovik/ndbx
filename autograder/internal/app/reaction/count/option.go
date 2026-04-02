package count

// Option represents a functional option for configuring the Counts instance.
type Option func(c *Counts)

// WithLikes sets the number of likes to the Counts instance.
func WithLikes(n uint64) Option {
	return func(c *Counts) {
		c.likes = n
	}
}

// WithDislikes sets the number of dislikes to the Counts instance.
func WithDislikes(n uint64) Option {
	return func(c *Counts) {
		c.dislikes = n
	}
}
