package graph

// Point represents a graph point identifier.
type Point struct {
	id string
}

// NewPoint creates a point with the provided identifier.
func NewPoint(id string) Point {
	return Point{
		id,
	}
}
