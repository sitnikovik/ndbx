package expectation

// Option is a functional option
// to configure the Expectations instance on creation.
type Option func(e *Expectations)

// WithCount sets the number of events
// to expect to the instance on creation.
func WithCount(n int) Option {
	return func(e *Expectations) {
		e.count = n
	}
}
