package times

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertEquals checks if the expected and actual times are equal (ignoring time zones).
//
// If the times are not equal, it returns an error indicating that the expectation has failed.
func AssertEquals(expected, actual time.Time) error {
	if !expected.UTC().Equal(actual.UTC()) {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected %s but got %s",
			log.Time(expected),
			log.Time(actual),
		)
	}
	return nil
}
