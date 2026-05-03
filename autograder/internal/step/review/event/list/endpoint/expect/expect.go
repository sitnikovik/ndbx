package expect

// Expectations holds the expectations we need to check in the Step.
type Expectations struct {
	// count is the expected count of reviews in response to expect.
	count int
}

// NewExpectations creates a new Expectations instance.
func NewExpectations(opt Option, opts ...Option) Expectations {
	e := Expectations{}
	opt(&e)
	for _, o := range opts {
		o(&e)
	}
	return e
}

// Count returns the expected count of reviews in response to expect.
func (e Expectations) Count() int {
	return e.count
}
