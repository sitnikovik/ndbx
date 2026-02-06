package numbers

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// IntegerEquality represents an equality check between two integers.
type IntegerEquality struct {
	// expected is the expected integer value.
	expected int
	// actual is the actual integer value.
	actual int
}

// NewIntegerEquality creates a new IntegerEquality instance.
//
// Parameters:
//   - expected: The expected integer value.
//   - actual: The actual integer value.
func NewIntegerEquality(expected, actual int) IntegerEquality {
	return IntegerEquality{
		expected: expected,
		actual:   actual,
	}
}

// Error checks if the expected and actual integers are equal.
//
// Returns an error if they are not equal, otherwise nil.
func (e IntegerEquality) Error() error {
	if e.expected != e.actual {
		return errs.Wrap(
			expect.ErrEqualityFailed,
			"want %s, got %s",
			log.Number(e.expected),
			log.Number(e.actual),
		)
	}
	return nil
}
