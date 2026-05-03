package expectation

import (
	"net/http"

	cookie "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
)

// Expectations holds the expectations we need to check in the http step.
type Expectations struct {
	// cookie is the cookie expectations.
	cookies []cookie.Expectations
	// asserts is the list of response assertions.
	asserts []response.AssertFunc
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

// Assert checks that the provided response satisfies all expectations
// and returns an error if any expectation fails, suitable for test validation.
func (e Expectations) Assert(resp *http.Response) error {
	var err error
	for _, asrt := range e.asserts {
		err = asrt(resp)
		if err != nil {
			return err
		}
	}
	for _, ck := range e.cookies {
		err = ck.Assert(resp.Cookies())
		if err != nil {
			return err
		}
	}
	return nil
}
