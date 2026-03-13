package times

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertNotAfter checks if the expected time is not after the actual time (ignoring time zones).
//
// If the expected time is after the actual time, it returns an error indicating that the expectation has failed.
func AssertNotAfter(expected, actual time.Time) error {
	if expected.UTC().After(actual.UTC()) {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected %s to not be after %s",
			log.Time(expected),
			log.Time(actual),
		)
	}
	return nil
}
