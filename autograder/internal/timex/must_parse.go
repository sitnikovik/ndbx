package timex

import "time"

// MustParse is a helper function that parses a time string according to the specified layout.
//
// Panics if the parsing fails, making it useful for situations where the time format is known to be correct.
func MustParse(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}
