package duration

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertEquals checks if the expected and actual durations are equal.
//
// If the durations are not equal, it returns an error indicating that the expectation has failed.
func AssertEquals(expected, actual time.Duration) error {
	if expected != actual {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"want %s but got %s",
			log.Duration(expected),
			log.Duration(actual),
		)
	}
	return nil
}
