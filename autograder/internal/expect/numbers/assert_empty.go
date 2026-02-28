package numbers

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// AssertEmpty checks if the given number is empty (zero) and returns an error if it is not.
func AssertEmpty[T number](num T) error {
	if num != 0 {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"got %s",
			log.Number(num),
		)
	}
	return nil
}
