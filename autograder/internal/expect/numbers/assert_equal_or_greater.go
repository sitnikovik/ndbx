package numbers

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertEqualOrGreater checks if the expected and actual integers
// are equal or the actual one is greater than expected
// and returns an error if not.
func AssertEqualOrGreater[T number](expected, actual T) error {
	if actual < expected {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expect '%s' to be equal or greater than '%s'",
			log.Number(actual),
			log.Number(expected),
		)
	}
	return nil
}
