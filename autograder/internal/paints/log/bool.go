package log

import (
	"fmt"

	"github.com/sitnikovik/paints/color/text"
)

// Bool returns a string representation of a boolean formatted for console logs.
func Bool(b bool) string {
	return text.Purple(fmt.Sprintf("%t", b))
}
