package count

// Counts represents counters of reactions for the objects.
type Counts struct {
	// likes is the number of likes users left.
	likes uint64
	// dislikes is the number of dislikes users left.
	dislikes uint64
}

// NewCounts creates a new Counts instance.
func NewCounts(opts ...Option) Counts {
	c := Counts{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// With copies the instance and returns a new one with provided options.
func (c Counts) With(opts ...Option) Counts {
	cop := c
	for _, opt := range opts {
		opt(&cop)
	}
	return cop
}

// Empty defines whether the Counts is empty.
func (c Counts) Empty() bool {
	return c.likes == 0 && c.dislikes == 0
}

// Likes returns the number of likes users left.
func (c Counts) Likes() uint64 {
	return c.likes
}

// Dislikes returns the number of dislikes users left.
func (c Counts) Dislikes() uint64 {
	return c.dislikes
}
