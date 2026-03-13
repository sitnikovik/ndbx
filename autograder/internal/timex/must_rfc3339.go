package timex

import "time"

// MustRFC3339 parses a time string in RFC3339 format.
//
// Panics if the parsing fails, making it useful for situations where the time format is known to be correct.
func MustRFC3339(tim string) time.Time {
	t, err := time.Parse(time.RFC3339, tim)
	if err != nil {
		panic(err)
	}
	return t
}
