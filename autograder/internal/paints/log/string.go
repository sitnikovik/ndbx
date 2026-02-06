package log

import "github.com/sitnikovik/paints/color/text"

// String returns a string representation of a string formatted for console logs.
func String(s string) string {
	return text.Yellow("\"" + s + "\"")
}
