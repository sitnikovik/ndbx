package expectation

// Expectations holds the expectations to verify in the Step.
type Expectations struct {
	// count is the number of rows to expect.
	count int
}

// NewExpectations creates a new Expectations instance with the given options.
func NewExpectations(opt Option, opts ...Option) Expectations {
	e := Expectations{}
	opt(&e)
	for _, o := range opts {
		o(&e)
	}
	return e
}

// Count returns the number of rows to expect.
func (e Expectations) Count() int {
	return e.count
}
