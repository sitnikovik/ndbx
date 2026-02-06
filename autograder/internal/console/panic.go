package console

import (
	"fmt"

	bg "github.com/sitnikovik/paints/color/background"
)

// Panic logs a formatted error message to the console but does not exit the program.
//
// Parameters:
//   - v: The panic value to be logged as a panic message.
//   - args: A variadic list of additional arguments to be included in the log message.
//
// Example usage:
//
//	Panic("nil pointer dereference", "foo", 1) // logs "panic: nil pointer dereference\nadditional info:\nfoo1"
func Panic(v any, args ...any) {
	f := bg.Red("panic") + ": %v"
	if len(args) > 0 {
		f += "\nadditional info:\n" + fmt.Sprint(args...)
	}
	Log(f, v)
}
