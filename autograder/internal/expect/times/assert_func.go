package times

import "time"

// AssertFunc is a function type that defines the signature
// for assertion functions that compare expected and actual time values.
type AssertFunc func(expected, actual time.Time) error
