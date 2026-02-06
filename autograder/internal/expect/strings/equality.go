package strings

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// StringEquality represents an equality check between two strings.
type StringEquality struct {
	// expected is the expected string value.
	expected string
	// actual is the actual string value.
	actual string
}

// NewStringEquality creates a new StringEquality instance.
//
// Parameters:
//   - expected: The expected string value.
//   - actual: The actual string value.
func NewStringEquality(expected, actual string) StringEquality {
	return StringEquality{
		expected: expected,
		actual:   actual,
	}
}

// Error checks if the expected and actual strings are equal.
//
// Returns an error if they are not equal, otherwise nil.
func (e StringEquality) Error() error {
	if e.expected != e.actual {
		return errs.Wrap(
			expect.ErrEqualityFailed,
			"want %s, got %s",
			log.String(e.expected),
			log.String(e.actual),
		)
	}
	return nil
}
