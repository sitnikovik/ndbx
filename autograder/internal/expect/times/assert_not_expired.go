package times

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertNotExpired checks if the time since the given time is less
// than the expected duration plus an actual duration.
//
// If the time since the given time exceeds this threshold,
// it returns an error indicating that the expectation has failed.
func AssertNotExpired(
	since time.Time,
	expected time.Duration,
	actual time.Duration,
) error {
	if time.Since(since) > expected+actual {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected not to be expired but it is (since %s, expected %s, actual %s)",
			log.Time(since),
			log.Duration(expected),
			log.Duration(actual),
		)
	}
	return nil
}
