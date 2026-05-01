package expectation

import (
	cookie "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
)

// Option is an functional option for the Expectations instance.
type Option func(*Expectations)

// WithAsserts is an option to set the response assertions on creation.
func WithAsserts(ff ...response.AssertFunc) Option {
	return func(e *Expectations) {
		e.asserts = ff
	}
}

// WithCookies is an option to set the cookies expectations on creation.
func WithCookies(ee ...cookie.Expectations) Option {
	return func(e *Expectations) {
		e.cookies = ee
	}
}
