package variable

import (
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Values is a wrapper around step.Variables that provides convenient methods
// to access specific variables related to user sessions.
type Values struct {
	// orig is the original instance that contains all the variables.
	orig step.Variables
}

// NewValues creates a new Values instance from the given step.Variables.
func NewValues(orig step.Variables) Values {
	return Values{
		orig: orig,
	}
}
