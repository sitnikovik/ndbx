package expect

// Option represents the functional option
// to configure the Exectations instance on its creation.
type Option func(e *Expectations)

// WithCount sets the expected count of reviews in response to expect.
func WithCount(count int) Option {
	return func(e *Expectations) {
		e.count = count
	}
}
