package errs

import (
	"fmt"
	"strings"
)

// WrapNested wraps an existing error with additional context, preserving the original error's message.
//
// Return an error that wraps the current error with additional context while preserving the previous error's message.
//
// If either the current error or the previous error is nil, it returns nil.
//
// The function takes the current error to wrap, the previous error whose message should be preserved,
// a format string for additional context, and variadic arguments for the format string.
// It returns an error that combines the current error
// with the additional context while preserving the message of the previous error.
//
// Parameters:
//   - curr: The current error to wrap.
//   - prev: The previous error whose message should be preserved.
//   - format: A format string for additional context.
//   - args: Variadic arguments for the format string.
//
// Example usage:
//
//	err1 := fmt.Errorf(context.DeadlineExceeded, "operation timed out")
//	err2 := WrapNested(context.DeadlineExceeded, err1, "additional context")
//	fmt.Println(err2) // context.DeadlineExceeded: additional context: operation timed out
func WrapNested(curr, prev error, format string, args ...any) error {
	if curr == nil || prev == nil {
		return nil
	}
	return fmt.Errorf(
		"%w: %s: %s",
		curr,
		fmt.Sprintf(format, args...),
		strings.TrimLeft(
			prev.Error(),
			curr.Error()+": ",
		),
	)
}
