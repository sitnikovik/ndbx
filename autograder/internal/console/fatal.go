package console

import (
	"fmt"
	"os"
)

// Fatal logs a formatted error message to the console and exits the program.
//
// Parameters:
//   - format: A format string.
//   - args: A variadic list of arguments to be formatted according to the format string.
//
// Example usage:
//
//	Fatal("Failed to open file: %s", err)
func Fatal(format string, args ...any) {
	fmt.Printf("❌ %s\n", fmt.Sprintf(format, args...))
	os.Exit(1)
}
