package numbers

// number is a type constraint that allows int, int64, and float64 types.
type number interface {
	int | int64 | float64
}
