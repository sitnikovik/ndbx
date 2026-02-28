package log

import "time"

// Time returns a string representation of time formatted for console logs
// and formats it using the RFC3339 format.
func Time(t time.Time) string {
	return String(t.Format(time.RFC3339))
}
