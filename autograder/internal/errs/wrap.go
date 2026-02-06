package errs

import "fmt"

// Wrap wraps an existing error with additional context.
//
// If the provided error is nil, it returns nil.
//
// Parameters:
//   - err: The original error to wrap.
//   - format: A format string for additional context.
//   - args: Variadic arguments for the format string.
//
// Returns:
//   - An error that wraps the original error with additional context.
func Wrap(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s", err, fmt.Sprintf(format, args...))
}
