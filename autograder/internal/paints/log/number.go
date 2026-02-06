package log

import (
	"fmt"

	"github.com/sitnikovik/paints/color/text"
)

// Number returns a string representation of an integer formatted for console logs.
func Number(n int) string {
	return text.Cyan(fmt.Sprintf("%d", n))
}
