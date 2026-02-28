package log

import (
	"time"
)

// Duration returns a string representation of a time.Duration formatted for console logs.
func Duration(d time.Duration) string {
	return String(d.String())
}
