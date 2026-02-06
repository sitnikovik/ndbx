package console

import "fmt"

// Log logs a formatted message to the console.
//
// Parameters:
//   - format: A format string.
//   - args: A variadic list of arguments to be formatted according to the format string.
//
// Example usage:
//
//	Log("Hello, %s!", "world") // Output: Hello, world!
func Log(format string, args ...any) {
	fmt.Printf("%s\n", fmt.Sprintf(format, args...))
}
