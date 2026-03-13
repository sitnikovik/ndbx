package response

import "net/http"

// AssertFunc is a function type that takes an HTTP response and returns an error if the assertion fails.
type AssertFunc func(rsp *http.Response) error

// AssertAll takes an HTTP response and a variadic number of assertion functions
// and returns an error if any of the assertions fail.
//
// Panics if no assertion functions are provided.
func AssertAll(rsp *http.Response, ff ...AssertFunc) error {
	if len(ff) == 0 {
		panic("at least one assertion function must be provided")
	}
	for _, f := range ff {
		if err := f(rsp); err != nil {
			return err
		}
	}
	return nil
}
