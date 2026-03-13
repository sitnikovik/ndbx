package strings

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertNotEquals checks if the expected and actual strings are not equal and returns an error if they are.
func AssertNotEquals(expected, actual string) error {
	if expected == actual {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"want %s, got %s",
			log.String(expected),
			log.String(actual),
		)
	}
	return nil
}
