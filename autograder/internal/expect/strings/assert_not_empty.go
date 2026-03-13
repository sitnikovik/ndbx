package strings

import "github.com/sitnikovik/ndbx/autograder/internal/errs"

// AssertNotEmpty checks if the given string is not empty and returns an error if it is.
func AssertNotEmpty(s string) error {
	if s == "" {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected non-empty string, but got the empty one",
		)
	}
	return nil
}
