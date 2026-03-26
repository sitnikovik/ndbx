package expect

import (
	"reflect"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// AssertEquals compares to objects of any type
// and returns an error if the one not equals to the second one.
func AssertEquals(expect, actual any) error {
	if !reflect.DeepEqual(expect, actual) {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"objects are not the same with data",
		)
	}
	return nil
}
