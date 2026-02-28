package numbers

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertEquals checks if the expected and actual integers are equal and returns an error if they are not.
func AssertEquals[T number](expected, actual T) error {
	if expected != actual {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"want %s, got %s",
			log.Number(expected),
			log.Number(actual),
		)
	}
	return nil
}
