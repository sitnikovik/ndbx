package log

import "github.com/sitnikovik/paints/color/text"

// URL returns a string representation of a URL formatted for console logs.
func URL(u string) string {
	return text.Green("\"" + u + "\"")
}
