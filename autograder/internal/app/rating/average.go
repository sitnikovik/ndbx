package rating

import "math"

// Average calculates the average rating of the given ratings,
// rounds the result to one decimal place (tenths),
// and returns a new Rating.
//
// Panics if any Rating is invalid.
func Average(rr ...Rating) Rating {
	if len(rr) == 0 {
		return NewRating(0.0)
	}
	var sum float64
	for _, r := range rr {
		sum += r.Exact()
	}
	avg := sum / float64(len(rr))
	return NewRating(math.Round(avg*10) / 10)
}
