package expectation

// Option is an functional option for the Expectations instance.
type Option func(*Expectations)

// WithAsserts is an option to set the cookie assertions on creation.
func WithAsserts(ff ...AssertFunc) Option {
	return func(e *Expectations) {
		e.asserts = ff
	}
}

// WithAssertsValueFn is an option to set the cookie value assertion on creation.
func WithAssertsValueFn(f AssertValueFunc) Option {
	return func(e *Expectations) {
		e.assertsValueFn = f
	}
}
