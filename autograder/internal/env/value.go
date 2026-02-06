package env

import (
	"fmt"
	"os"

	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

// Value represents the value of an environment variable.
type Value string

// NewValue creates a new Value instance.
//
// Parameters:
//   - v: The string value of the environment variable.
func NewValue(v string) Value {
	return Value(v)
}

// Get retrieves the value of the specified environment variable.
//
// Parameters:
//   - name: The name of the environment variable.
func Get(name string) Value {
	return Value(os.Getenv(name))
}

// MustGet retrieves the value of the specified environment variable
// and panics if it is not set.
//
// Parameters:
//   - name: The name of the environment variable.
func MustGet(name string) Value {
	v := Get(name)
	if v.Empty() {
		panic(fmt.Sprintf(
			"environment variable %s is required",
			log.String(name),
		))
	}
	return v
}

// Empty checks if the environment variable value is empty.
func (v Value) Empty() bool {
	return v == ""
}

// String returns the string representation of the environment variable value.
func (v Value) String() string {
	return string(v)
}
