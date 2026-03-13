package times

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertAfterOrEqual checks if the actual time is after
// or equal to the expected time (ignoring time zones).
func AssertAfterOrEqual(
	actual time.Time,
	expected time.Time,
) error {
	t1 := actual.UTC()
	t2 := expected.UTC()
	if t2.After(t1) || t2.Equal(t1) {
		return nil
	}
	return errs.Wrap(
		errs.ErrExpectationFailed,
		"expect %s to be after or equal to %s",
		log.Time(actual),
		log.Time(expected),
	)
}
