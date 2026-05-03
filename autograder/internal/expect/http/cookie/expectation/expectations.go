package expectation

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie"
)

// AssertFunc asserts cookie.
type AssertFunc func(ck []*http.Cookie, name string) error

// AssertValueFunc asserts cookie value.
type AssertValueFunc func(v string) error

// Expectations is a set of expectations for cookies.
type Expectations struct {
	// asserts is a list of asserts to run.
	asserts []AssertFunc
	// assertsValueFn is a function to assert cookie value.
	assertsValueFn AssertValueFunc
	// name is a name of the cookie to check.
	name string
}

// NewExpectations creates new Expectations instance
// with the target cookie name and options.
func NewExpectations(
	name string,
	opt Option,
	opts ...Option,
) Expectations {
	e := Expectations{
		name: name,
	}
	opt(&e)
	for _, o := range opts {
		o(&e)
	}
	return e
}

// Assert checks that the provided cookies satisfy all expectations
// and returns an error if any expectation fails, suitable for test validation.
func (e Expectations) Assert(ckk []*http.Cookie) error {
	var err error
	if e.assertsValueFn != nil {
		err = cookie.AssertValueFn(ckk, e.name, e.assertsValueFn)
		if err != nil {
			return err
		}
	}
	for _, asrt := range e.asserts {
		err = asrt(ckk, e.name)
		if err != nil {
			return err
		}
	}
	return nil
}
