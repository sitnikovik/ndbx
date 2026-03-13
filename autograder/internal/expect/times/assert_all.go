package times

import "time"

// AssertAll applies multiple assertion functions to the expected and actual time values.
//
// It returns an error if any of the assertions fail. If no assertion functions are provided,
// it panics to indicate that at least one assertion function must be supplied.
func AssertAll(
	expected,
	actual time.Time,
	ff ...AssertFunc,
) error {
	if len(ff) == 0 {
		panic("at least one assertion function must be provided")
	}
	for _, f := range ff {
		if err := f(expected, actual); err != nil {
			return err
		}
	}
	return nil
}
