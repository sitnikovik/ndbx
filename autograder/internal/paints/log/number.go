package log

import (
	"fmt"

	"github.com/sitnikovik/paints/color/text"
)

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Number returns a string representation of an integer formatted for console logs.
func Number[T number](n T) string {
	return text.Cyan(fmt.Sprintf("%v", n))
}
