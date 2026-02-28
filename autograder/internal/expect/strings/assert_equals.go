package strings

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertEquals checks if the expected and actual strings are equal and returns an error if they are not.
func AssertEquals(expected, actual string) error {
	if expected != actual {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"want %s, got %s",
			log.String(expected),
			log.String(actual),
		)
	}
	return nil
}
